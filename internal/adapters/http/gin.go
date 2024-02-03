package http

import (
	"fmt"
	"net/http"

	"github.com/AntonyIS/notelify-logging-service/config"
	"github.com/AntonyIS/notelify-logging-service/internal/core/domain"
	"github.com/AntonyIS/notelify-logging-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginHandler interface {
	PostLog(ctx *gin.Context)
	GetLogs(ctx *gin.Context)
	GetServiceLogs(ctx *gin.Context)
	GetServiceLogsByLogLevel(ctx *gin.Context)
	HealthCheck(ctx *gin.Context)
}

type handler struct {
	svc ports.LoggerService
}

func NewGinHandler(svc ports.LoggerService) ginHandler {
	router := handler{
		svc: svc,
	}
	return router
}

func (h handler) PostLog(ctx *gin.Context) {
	var logEntry domain.LogMessage
	if err := ctx.ShouldBindJSON(&logEntry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	go h.svc.CreateLog(logEntry)
	ctx.JSON(http.StatusCreated, gin.H{"message": "message posted successfuly"})
}

func (h handler) GetLogs(ctx *gin.Context) {
	response := h.svc.GetLogs()
	ctx.JSON(http.StatusOK, response)

}
func (h handler) GetServiceLogs(ctx *gin.Context) {
	service := ctx.Param("service")
	response := h.svc.GetServiceLogs(service)
	ctx.JSON(http.StatusOK, response)
}

func (h handler) GetServiceLogsByLogLevel(ctx *gin.Context) {
	service := ctx.Param("service")
	log_level := ctx.Param("log_level")
	response := h.svc.GetServiceLogsByLogLevel(service, log_level)
	ctx.JSON(http.StatusOK, response)

}

func (h handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Server running", "status": "success"})
}

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
		
		logRoutes.GET("/", logHandler.GetLogs)
		logRoutes.GET("/healthcheck", logHandler.HealthCheck)
		logRoutes.GET("/:service", logHandler.GetServiceLogs)
		logRoutes.GET("/:service/:log_level", logHandler.GetServiceLogsByLogLevel)
		logRoutes.POST("/:service", logHandler.PostLog)
	}
	router.Run(fmt.Sprintf(":%s", conf.SERVER_PORT))
}
