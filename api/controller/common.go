package controller

import (
	"example/komposervice/internal/schema"
	"example/komposervice/internal/service"

	"github.com/gin-gonic/gin"
)

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
