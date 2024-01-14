package domain

type LogMessage struct {
	Log_id  string `json:"log_id"`
	Message string `json:"message"`
	Service string `json:"service"`
}
