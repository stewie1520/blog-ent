package usecases

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrEntityNotFound = errors.New("entity not found")
var ErrEntityForbidden = errors.New("forbidden")
