package ports

import "github.com/AntonyIS/notelify-logging-svc/internal/core/domain"

type LoggerService interface {
	Log(message, service string)
}

type LoggerRepository interface {
	Log(message domain.LogMessage)
}
