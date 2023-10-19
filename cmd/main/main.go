package main

import (
	"flag"
	"fmt"

	"github.com/stewie1520/blog_ent/api"
	"github.com/stewie1520/blog_ent/config"
	"github.com/stewie1520/blog_ent/core"
	docs "github.com/stewie1520/blog_ent/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var debug = flag.Bool("debug", false, "debug mode")

func init() {
	flag.Parse()
}

// @title Blog API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	cfg, err := config.Init()
	panicIfError(err)

	app := core.NewBaseApp(core.BaseAppConfig{
		Config:  cfg,
		IsDebug: *debug,
	})

	err = app.Bootstrap()
	panicIfError(err)

	router, err := api.InitApi(app)
	panicIfError(err)

	if !app.Config().IsProd() {
		docs.SwaggerInfo.BasePath = "/"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	err = router.Run(fmt.Sprintf(":%d", cfg.Port))
	panicIfError(err)
}

func panicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
