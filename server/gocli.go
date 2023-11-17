package main

import (
	"database/sql"
	"fmt"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/CLI"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	level            string
	message          string
	resourceID       string
	timestamp        string
	traceID          string
	spanID           string
	commit           string
	parentResourceID string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "query-processor",
	Short: "A simple log query processor made for dyte",
	Run:   runQuery,
}

func init() {
	rootCmd.Flags().StringVarP(&level, "level", "l", "", "Filter logs by level")
	rootCmd.Flags().StringVarP(&message, "message", "m", "", "Filter logs by message")
	rootCmd.Flags().StringVarP(&resourceID, "resourceId", "r", "", "Filter logs by resource ID")
	rootCmd.Flags().StringVarP(&timestamp, "timestamp", "t", "", "Filter logs by timestamp")
	rootCmd.Flags().StringVarP(&traceID, "TraceId", "T", "", "Filter logs by Trace ID")
	rootCmd.Flags().StringVarP(&spanID, "spanId", "s", "", "Filter logs by Span ID")
	rootCmd.Flags().StringVarP(&commit, "Commit", "c", "", "Filter logs by Commit")
	rootCmd.Flags().StringVarP(&parentResourceID, "ParentResId", "p", "", "Filter logs by Parent Resource ID")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runQuery(cmd *cobra.Command, args []string) {
	// Establish a connection to the MySQL database
	var logger *log.Logger
	logger = comms.LoggerInit()
	db, err := sql.Open("mysql", "admin:GgDA7yFpfM2We2@tcp(database-level-error.czzptrur4kwd.eu-north-1.rds.amazonaws.com)/Logger")

	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal("Unable to close the database connection")
		}
	}(db)

	// Construct the SQL query based on the provided flags
	query, argsList := CLI.PrepareGeneralQuery(level, message, resourceID, timestamp, traceID, spanID, commit, parentResourceID)
	// Add other conditions as needed
	logger.Println("creating the prepare statement")

	stmt, err := CLI.PrepareGeneralStatement(db, query)
	if err != nil {
		log.Fatal("failed to prepare statements", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.Fatal("failed to close the prepared statement")
		}
	}(stmt)

	// Execute the query
	rows, err := stmt.Query(argsList...)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	CLI.PrintRows(rows)
}
