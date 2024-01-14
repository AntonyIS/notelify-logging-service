package cmd

import (
	"github.com/AntonyIS/notelify-logging-svc/config"
	"github.com/AntonyIS/notelify-logging-svc/internal/adapters/http"
	"github.com/AntonyIS/notelify-logging-svc/internal/adapters/repository/postgres"
	"github.com/AntonyIS/notelify-logging-svc/internal/core/services"
)

func RunService() {
	// Read application environment and load configurations
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	databaseRepo, _ := postgres.NewPostgresClient(*conf)
	loggerSvc := services.NewLoggingManagementService(databaseRepo)
	http.InitGinRoutes(loggerSvc, *conf)
}
