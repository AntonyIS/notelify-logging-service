package ports

import "github.com/AntonyIS/notelify-logging-service/internal/core/domain"

type LoggerService interface {
	CreateLog(message, service string)
	GetLogs() *[]domain.LogMessage
}

type LoggerRepository interface {
	CreateLog(message domain.LogMessage)
	GetLogs() *[]domain.LogMessage
}
