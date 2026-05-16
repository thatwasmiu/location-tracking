package router

import (
	"location-tracking/internal/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthHlr := handler.NewHealthHandler()

	api := router.Group("/api")

	{
		api.GET("/health", healthHlr.Health)
	}

	return router
}
