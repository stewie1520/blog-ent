package middleware

import (
	ginCors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog_ent/core"
)

func Cors(app core.App) gin.HandlerFunc {
	return ginCors.New(ginCors.Config{
		AllowOrigins:     []string{app.Config().WebsiteDomain},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowCredentials: true,
	})
}
