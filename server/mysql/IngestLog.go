package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/model"
	"log"
)

func IngestLogs(logEntry model.LogEntry, pool *DBPool, logger *log.Logger) error {
	tableName := comms.GetTargetTable(logEntry.Level)
	logger.Println(fmt.Sprintf("we are chossing the %s shard", tableName))
	stmt, err := pool.Prepare("INSERT INTO " + string(tableName) + " (level, message, resourceId, timestamp, traceId, spanId, commit, metadata_parentResourceId) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		//return fmt.Errorf("error preparing SQL statement: %v", err)
		logger.Println("failed to prepare statements", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.Fatal("Failed to close the prepare statements")
		}
	}(stmt)
	_, err = stmt.Exec(logEntry.Level, logEntry.Message, logEntry.ResourceID, logEntry.Timestamp, logEntry.TraceID, logEntry.SpanID, logEntry.Commit, logEntry.Metadata.ParentResourceID)
	if err != nil {
		logger.Println("error executing SQL statement: %v", err)
		return err
	}
	return nil
}
