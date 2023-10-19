package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog_ent/api/middleware"
	"github.com/stewie1520/blog_ent/api/response"
	"github.com/stewie1520/blog_ent/core"
	"github.com/stewie1520/blog_ent/usecases"
	usecases_user "github.com/stewie1520/blog_ent/usecases/user"
)

type userApi struct {
	app core.App
}

func bindUserApi(app core.App, ginEngine *gin.Engine) {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/users")
	subGroup.POST("/register", api.register)
	subGroup.POST("/login", api.login)

	subGroup.Use(middleware.RequireAuth(app))

	subGroup.GET("/me", api.me)
}

// register Register new user
// @Summary Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body usecases_user.RegisterCommand true "Register payload"
// @Success 200 {object} usecases_user.TokensResponse
// @Router /users/register [post]
func (api *userApi) register(c *gin.Context) {
	cmd := usecases_user.NewRegisterCommand(api.app)

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	if res, err := cmd.Execute(c); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

// login Login
// @Summary Login
// @Tags user
// @Accept json
// @Produce json
// @Param user body usecases_user.LoginCommand true "Login payload"
// @Success 200 {object} usecases_user.TokensResponse
// @Router /users/login [post]
func (api *userApi) login(c *gin.Context) {
	cmd := usecases_user.NewLoginCommand(api.app)

	if err := c.ShouldBindJSON(cmd); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	res, err := cmd.Execute(c)

	if err == usecases.ErrInvalidCredentials {
		response.NewBadRequestError("Failed to authenticate", err).WithGin(c)
		return
	}

	if err != nil {
		response.NewBadRequestError("", err).WithGin(c)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// me Get current user
// @Summary Get current user
// @Tags user
// @Accept json
// @Produce json
// @Success 200
// @Router /users/me [get]
// @Security Authorization
func (api *userApi) me(c *gin.Context) {
	q := usecases_user.NewGetUserQuery(api.app)
	q.ID = getUserFromContext(c).ID.String()

	if res, err := q.Execute(c); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}
