package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "Ping success"})

}
