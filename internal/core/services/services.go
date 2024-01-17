package services

import (
	"github.com/AntonyIS/notelify-logging-service/internal/core/domain"
	"github.com/AntonyIS/notelify-logging-service/internal/core/ports"
	"github.com/google/uuid"
)

type loggingManagementService struct {
	repo ports.LoggerRepository
}

func NewLoggingManagementService(repo ports.LoggerRepository) *loggingManagementService {
	svc := loggingManagementService{
		repo: repo,
	}
	return &svc
}

func (svc *loggingManagementService) CreateLog(logEntry domain.LogMessage) {
	logEntry.LogID = uuid.New().String()
	err := svc.repo.CreateLog(logEntry)
	if err != nil {
		var logEntry domain.LogMessage
		logEntry.Message = err.Error()
		logEntry.LogLevel = "Error"
		logEntry.Service = "Logger"
		svc.repo.CreateLog(logEntry)
	}
}

func (svc *loggingManagementService) GetLogs() *[]domain.LogMessage {
	logs, err := svc.repo.GetLogs()
	if err != nil {
		var logEntry domain.LogMessage
		logEntry.Message = err.Error()
		logEntry.LogLevel = "Error"
		logEntry.Service = "Logger"
		svc.repo.CreateLog(logEntry)
	}
	return logs
}

func (svc *loggingManagementService) GetServiceLogs(service string) *[]domain.LogMessage {
	logs, err := svc.repo.GetServiceLogs(service)
	if err != nil {
		var logEntry domain.LogMessage
		logEntry.Message = err.Error()
		logEntry.LogLevel = "Error"
		logEntry.Service = "Logger"
		svc.repo.CreateLog(logEntry)
	}
	return logs
}

func (svc *loggingManagementService) GetServiceLogsByLogLevel(service, log_level string) *[]domain.LogMessage {
	logs, err := svc.repo.GetServiceLogsByLogLevel(service, log_level)
	if err != nil {
		var logEntry domain.LogMessage
		logEntry.Message = err.Error()
		logEntry.LogLevel = "Error"
		logEntry.Service = "Logger"
		svc.repo.CreateLog(logEntry)
	}
	return logs
}
