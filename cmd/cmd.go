package cmd

import (
	"fmt"

	"github.com/AntonyIS/notelify-logging-service/config"
	"github.com/AntonyIS/notelify-logging-service/internal/adapters/http"
	"github.com/AntonyIS/notelify-logging-service/internal/adapters/repository/postgres"
	"github.com/AntonyIS/notelify-logging-service/internal/core/services"
)

func RunService() {
	// Read application environment and load configurations
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	repo, err := postgres.NewPostgresClient(*conf)
	if err != nil {
		panic(err)
	}
	loggerSvc := services.NewLoggingManagementService(repo)
	fmt.Println(conf, loggerSvc.GetLogs())
	http.InitGinRoutes(loggerSvc, *conf)
}
