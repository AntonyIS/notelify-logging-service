package http

import (
	"fmt"
	"net/http"

	"github.com/AntonyIS/notelify-logging-svc/config"
	"github.com/AntonyIS/notelify-logging-svc/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginHandler interface {
	PostLog(ctx *gin.Context)
	GetLog(ctx *gin.Context)
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
	var res string
	service := ctx.Param("service")
	if err := ctx.ShouldBindJSON(&res); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.svc.Log(res, service)
	ctx.JSON(http.StatusCreated, gin.H{"message": "message posted successfuly"})
}

func (h handler) GetLog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Logging service"})
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
	logRoutes := router.Group("/v1/logger")
	{
		logRoutes.POST("/:service", logHandler.PostLog)
		logRoutes.GET("/", logHandler.GetLog)
	}
	router.Run(fmt.Sprintf(":%s", conf.SERVER_PORT))
}
