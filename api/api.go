package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog_ent/api/middleware"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/ent"
)

func InitApi(app core.App) (*gin.Engine, error) {
	if app.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	bindHealthApi(app, engine)

	engine.Use(middleware.Cors(app))
	engine.Use(middleware.LoadAuthContext(app))

	bindUserApi(app, engine)

	return engine, nil
}

func getUserFromContext(c *gin.Context) *ent.User {
	value, _ := c.Get(middleware.ContextUserKey)
	user, ok := value.(*ent.User)

	if user == nil || !ok {
		return nil
	}

	return user
}
