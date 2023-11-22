package Models

import (
	"github.com/rohanailoni/Log-Ingressor/model"
)

type LogEntryWrapper struct {
	PrimaryKey int
	Logs       model.LogEntry
}
