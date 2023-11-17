package model

type LogEntry struct {
	Level      string `json:"level"`
	Message    string `json:"message"`
	ResourceID string `json:"resourceId"`
	Timestamp  string `json:"timestamp"`
	TraceID    string `json:"traceId"`
	SpanID     string `json:"spanId"`
	Commit     string `json:"commit"`
	Metadata   MetaData
}
