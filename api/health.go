package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/blog_ent/core"
)

type healthApi struct {
	app core.App
}

type ReadinessResponse struct {
	Healthy  bool   `json:"healthy"`
	Database string `json:"database,omitempty"` // won't be shown if empty
}

type LivenessResponse struct {
	Healthy bool `json:"healthy"`
}

func bindHealthApi(app core.App, ginEngine *gin.Engine) {
	api := &healthApi{
		app: app,
	}

	subGroup := ginEngine.Group("/health")
	subGroup.GET("/ready", api.readiness)
	subGroup.GET("/live", api.liveness)
}

// liveness
// @Summary Check if application is live
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} LivenessResponse
// @Router /health/live [get]
func (api *healthApi) liveness(c *gin.Context) {
	c.JSON(http.StatusOK, LivenessResponse{true})
}

// readiness
// @Summary Check if application is ready to serve traffic
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} ReadinessResponse
// @Router /health/ready [get]
func (api *healthApi) readiness(c *gin.Context) {
	c.JSON(http.StatusOK, ReadinessResponse{
		Healthy: true,
	})
}
