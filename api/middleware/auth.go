package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stewie1520/blog_ent/api/response"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/ent"
	"github.com/stewie1520/blog_ent/tools/securities"
)

const (
	ContextUserKey string = "user"
)

func LoadAuthContext(app core.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"]
		if len(authorization) == 0 {
			c.Next()
			return
		}

		token := authorization[0]

		if token == "" {
			c.Next()
			return
		}

		claims, err := securities.ParsePaseto(token, app.Config().Token.PublicKey)
		if err != nil {
			c.Next()
			return
		}

		userId, err := uuid.Parse(claims["userId"].(string))
		if err != nil {
			c.Next()
			return
		}

		user, err := app.Dao().User.Get(c, userId)
		if err != nil {
			c.Next()
			return
		}

		c.Set(ContextUserKey, &user)
	}
}

func RequireAuth(app core.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, _ := c.Get(ContextUserKey)
		user, ok := value.(*ent.User)

		if user == nil || !ok {
			response.NewUnauthorizedError("The request requires valid user authorization token to be set.", nil).WithGin(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
