package CLI

import (
	"database/sql"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
)

func PrintRows(rows *sql.Rows) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"PrimaryKey", "level", "message", "resourceId", "timestamp", "traceId", "spanID", "commit", "metadata_parentResourceId"})

	for rows.Next() {
		var logEntry model.LogEntry
		var key int
		err := rows.Scan(
			&key,
			&logEntry.Level,
			&logEntry.Message,
			&logEntry.ResourceID,
			&logEntry.Timestamp,
			&logEntry.TraceID,
			&logEntry.SpanID,
			&logEntry.Commit,
			&logEntry.Metadata.ParentResourceID,
		)
		if err != nil {
			log.Fatal(err)
		}
		t.AppendRow([]interface{}{key, logEntry.Level, logEntry.Message, logEntry.ResourceID, logEntry.Timestamp, logEntry.TraceID, logEntry.SpanID, logEntry.Commit, logEntry.Metadata.ParentResourceID})
	}
	t.Render()
}
