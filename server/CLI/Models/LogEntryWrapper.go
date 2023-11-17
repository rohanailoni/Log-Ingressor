package Models

import (
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/model"
)

type LogEntryWrapper struct {
	PrimaryKey int
	Logs       model.LogEntry
}
