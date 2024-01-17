package ports

import "github.com/AntonyIS/notelify-logging-service/internal/core/domain"

type LoggerService interface {
	CreateLog(logEntry domain.LogMessage)
	GetLogs() *[]domain.LogMessage
	GetServiceLogs(service string) *[]domain.LogMessage
	GetServiceLogsByLogLevel(service, log_level string) *[]domain.LogMessage
}

type LoggerRepository interface {
	CreateLog(logEntry domain.LogMessage) error
	GetLogs() (*[]domain.LogMessage, error)
	GetServiceLogs(service string) (*[]domain.LogMessage, error)
	GetServiceLogsByLogLevel(service, log_level string) (*[]domain.LogMessage, error)
}
