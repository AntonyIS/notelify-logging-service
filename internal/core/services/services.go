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

func (svc *loggingManagementService) CreateLog(messageString string, service string) {
	message := domain.LogMessage{
		Message: messageString,
		Log_id:  uuid.New().String(),
		Service: service,
	}
	svc.repo.CreateLog(message)
}

func (svc *loggingManagementService) GetLogs() *[]domain.LogMessage {
	return svc.repo.GetLogs()
}
