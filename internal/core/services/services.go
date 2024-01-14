package services

import (
	"github.com/AntonyIS/notelify-logging-svc/internal/core/domain"
	"github.com/AntonyIS/notelify-logging-svc/internal/core/ports"
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

func (svc *loggingManagementService) Log(messageString string, service string) {
	message := domain.LogMessage{
		Message: messageString,
		Log_id:  uuid.New().String(),
		Service: service,
	}
	svc.repo.Log(message)
}
