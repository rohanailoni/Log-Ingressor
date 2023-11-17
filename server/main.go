package main

//password:-GgDA7yFpfM2We2 //will not terminate the instance till 27-11-2023
import (
	comms "github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	handler "github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/handler"
	mysql "github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

func main() {
	//Initialisations.
	logger := comms.LoggerInit()
	config, err := comms.ReadConfig(comms.Configfile)
	if err != nil {
		logger.Println("Error reading the config file %s", err)
		return
	}
	pool, err := mysql.InitializeMySQLPool(&config)
	if err != nil {
		logger.Println("Error intialising the pool of connectins to mysql :%s", err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//Registering Handlers;
	r.GET("/", handler.PingHandler)
	r.POST("/", func(c *gin.Context) {
		handler.LogHandler(c, pool, logger)
	})

	//server configurations
	portEnv := os.Getenv(comms.ServerPortENV)
	if portEnv == "" {
		logger.Println("portenv is not set so falling back to config port if you are using docker please set environment variable of name ", comms.ServerPortENV)
		portEnv = config.Server.Address + ":" + strconv.Itoa(config.Server.Port)
	} else {
		portEnv = "0.0.0.0:" + portEnv
	}
	//starting the server.
	if err := r.Run(portEnv); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
