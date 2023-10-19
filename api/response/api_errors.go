// Reference: https://github.com/pocketbase/pocketbase/blob/master/apis/api_error.go

package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/stewie1520/blog_ent/log"
	"go.uber.org/zap"
)

type ApiError struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`

	rawData any
}

// NewApiError creates and returns new normalized `ApiError` instance.
func NewApiError(status int, message string, data any) *ApiError {
	return &ApiError{
		rawData: data,
		Data:    safeErrorsData(data),
		Code:    status,
		Message: strings.TrimSpace(message),
	}
}

// Error makes it compatible with the `error` interface.
func (e *ApiError) Error() string {
	return e.Message
}

// RawData returns the unformatted error data (could be an internal error, text, etc.)
func (e *ApiError) RawData() any {
	return e.rawData
}

func (e *ApiError) WithGin(c *gin.Context) {
	c.JSON(e.Code, e)
}

func (e *ApiError) WithResponseWriter(res http.ResponseWriter) {
	res.WriteHeader(e.Code)
	res.Header().Add("Content-Type", "application/json")
	result, _ := json.Marshal(e)

	_, err := res.Write(result)
	if err != nil {
		log.L().Info("Error writing response:", zap.Error(err))
	}
}

// NewBadRequestError creates and returns 400 `ApiError`.
func NewBadRequestError(message string, data any) *ApiError {
	if message == "" {
		message = "Something went wrong while processing your request."
	}

	return NewApiError(http.StatusBadRequest, message, data)
}

// NewUnauthorizedError creates and returns 401 `ApiError`.
func NewUnauthorizedError(message string, data any) *ApiError {
	if message == "" {
		message = "Missing or invalid authentication token."
	}

	return NewApiError(http.StatusUnauthorized, message, data)
}

func safeErrorsData(data any) map[string]any {
	switch v := data.(type) {
	case validation.Errors:
		return resolveSafeErrorsData[error](v)
	case map[string]validation.Error:
		return resolveSafeErrorsData[validation.Error](v)
	case map[string]error:
		return resolveSafeErrorsData[error](v)
	case map[string]any:
		return resolveSafeErrorsData[any](v)
	// gin use validator underlyings for binding errors
	case validator.ValidationErrors:
		return resolveSafeErrorsDataBinding(v)
	default:
		return map[string]any{} // not nil to ensure that is json serialized as object
	}
}

func resolveSafeErrorsData[T any](data map[string]T) map[string]any {
	result := map[string]any{}

	for name, err := range data {
		if isNestedError(err) {
			result[name] = safeErrorsData(err)
			continue
		}
		result[name] = resolveSafeErrorItem(err)
	}

	return result
}

func resolveSafeErrorsDataBinding(data validator.ValidationErrors) map[string]any {
	result := map[string]any{}

	for _, err := range data {
		result[err.Field()] = resolveSafeErrorItem(err)
	}

	return result
}

func isNestedError(err any) bool {
	switch err.(type) {
	case validation.Errors, map[string]validation.Error, map[string]error, map[string]any:
		return true
	}

	return false
}

// resolveSafeErrorItem extracts from each validation error its
// public safe error code and message.
func resolveSafeErrorItem(err any) map[string]string {
	// default public safe error values
	code := "validation_invalid_value"
	msg := "Invalid value."

	// only validation errors are public safe
	if obj, ok := err.(validation.Error); ok {
		code = obj.Code()
		msg = obj.Error()
	}

	if obj, ok := err.(validator.FieldError); ok {
		msg = fmt.Sprintf("field %s is invalid because of the %s tag", obj.Field(), obj.Tag())
	}

	return map[string]string{
		"code":    code,
		"message": msg,
	}
}
