package controller

import (
	"example/komposervice/internal/schema"
	"example/komposervice/internal/service"

	"github.com/gin-gonic/gin"
)

// HealthCheck.
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/common/healthcheck [GET]
func HealthCheck(c *gin.Context) {
	c.JSON(200, service.HealthCheck())
}

func WorkerCheck(c *gin.Context) {
	if err := service.WorkerCheck(); err != nil {
		c.JSON(400, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(200, service.HealthCheck())
}
