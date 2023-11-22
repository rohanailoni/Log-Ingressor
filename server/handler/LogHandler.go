package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanailoni/Log-Ingressor/comms"
	models "github.com/rohanailoni/Log-Ingressor/model"
	"github.com/rohanailoni/Log-Ingressor/mysql"
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
	context.JSON(http.StatusOK, gin.H{"status": "Log inserted successfully"})
}
