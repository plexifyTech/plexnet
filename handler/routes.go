package handler

import "github.com/gin-gonic/gin"

func RegisterAdminAPI(router gin.IRouter) {
	monitor := router.Group("")
	{
		monitor.GET("/health", HealthEndpoint)
	}
}
