package mysql

import (
	comms "github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	"strconv"
)

func InitializeMySQLPool(config *comms.Config) (*DBPool, error) {
	// Open a connection to the MySQL database
	//connectionString := "admin:GgDA7yFpfM2We2@tcp(database-level-error.czzptrur4kwd.eu-north-1.rds.amazonaws.com:3306)/"
	//db, err := sql.Open("mysql", connectionString)
	//if err != nil {
	//	log.Fatal("Error connecting to MySQL:", err)
	//}
	//
	//// Check the connection
	//if err := db.Ping(); err != nil {
	//	log.Fatal("Error pinging MySQL:", err)
	//}
	//
	//dbMutex.Lock()
	//db.SetMaxOpenConns(30) // Set maximum open connections
	//dbMutex.Unlock()
	//
	//fmt.Println("Connected to MySQL")
	db, err := NewDBPool(config.Database.User, config.Database.Password, config.Database.Host, strconv.Itoa(config.Database.Port), config.Database.Database, config.Database.MaxOpenConns, config.Database.MaxIdleConns)
	if err != nil {
		return nil, err
	}
	return db, nil
}
