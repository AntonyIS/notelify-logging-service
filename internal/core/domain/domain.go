package domain

type LogMessage struct {
	LogID   string `json:"log_id"`
	LogLevel string `json:"log_level"`
	Message  string `json:"message"`
	Service  string `json:"service"`
}
