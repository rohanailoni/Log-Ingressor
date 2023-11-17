package handler

import (
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	models "github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/model"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LogHandler(context *gin.Context, pool *mysql.DBPool, logger *log.Logger) {
	var logEntry models.LogEntry

	// Bind the JSON payload to the LogEntry struct
	if err := context.BindJSON(&logEntry); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	for i := 0; i < comms.RetiresIngest; i++ {
		logger.Println("Retry Ingest times ", i+1)
		err = mysql.IngestLogs(logEntry, pool, logger)
		if err == nil {
			if i > 0 {
				logger.Println("retries success after {}tries", i+1)
			}
			break
		}
	}
	if err != nil {
		logger.Fatal("Failed to execute the logger with log ingest %s", logEntry)
	}
	context.JSON(http.StatusOK, gin.H{"status": "Log inserted sucessfully"})
}
