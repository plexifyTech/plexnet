package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// HealthEndpoint HealthCheck
// @Summary Get the health status of the server
// @Description Get the health status of the  server.
// @Tags admin
// @Accept */*
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /admin/health [get]
func HealthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{"UP"})
}
