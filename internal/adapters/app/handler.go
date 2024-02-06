package app

import (
	"fmt"

	"github.com/AntonyIS/notelify-logging-service/config"
	"github.com/AntonyIS/notelify-logging-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(svc ports.LoggerService, conf config.Config) {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	logHandler := NewGinHandler(svc)
	logRoutes := router.Group("/logger/v1")
	{
		logRoutes.POST("/:service", logHandler.PostLog)
		logRoutes.GET("/", logHandler.GetLogs)
		logRoutes.GET("/healthcheck", logHandler.HealthCheck)
		logRoutes.GET("/:service", logHandler.GetServiceLogs)
		logRoutes.GET("/:service/:log_level", logHandler.GetServiceLogsByLogLevel)

	}
	router.Run(fmt.Sprintf(":%s", conf.SERVER_PORT))
}
