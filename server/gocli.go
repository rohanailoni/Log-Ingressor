package main

import (
	"database/sql"
	"fmt"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/CLI/Models"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jedib0t/go-pretty/v6/table"
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
	db, err := sql.Open("mysql", "admin:GgDA7yFpfM2We2@tcp(database-level-error.czzptrur4kwd.eu-north-1.rds.amazonaws.com)/Logger")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Construct the SQL query based on the provided flags
	query := "SELECT * FROM ErrorLog WHERE 1=1"
	argsList := []interface{}{}

	if level != "" {
		query += " AND level = ?"
		argsList = append(argsList, level)
	}

	if message != "" {
		query += " AND message = ?"
		argsList = append(argsList, message)
	}

	if resourceID != "" {
		query += " AND resourceId = ?"
		argsList = append(argsList, resourceID)
	}

	if timestamp != "" {
		query += " AND timestamp = ?"
		argsList = append(argsList, timestamp)
	}

	if traceID != "" {
		query += " AND TraceId = ?"
		argsList = append(argsList, traceID)
	}

	if spanID != "" {
		query += " AND spanId = ?"
		argsList = append(argsList, spanID)
	}

	if commit != "" {
		query += " AND Commit = ?"
		argsList = append(argsList, commit)
	}

	if parentResourceID != "" {
		query += " AND metadata_parentResourceId = ?"
		argsList = append(argsList, parentResourceID)
	}

	// Add other conditions as needed
	fmt.Println(query)
	// Execute the query
	rows, err := db.Query(query, argsList...)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Process the result set
	var logs []Models.LogEntryWrapper
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

	// Print the retrieved logs
	for _, logEntry := range logs {
		fmt.Printf("%+v\n", logEntry)
	}
	t.Render()
}
