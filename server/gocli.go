package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/CLI"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/CLI/Models"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/comms"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"log"
	"os"
)

//var (
//	level            string
//	message          string
//	resourceID       string
//	fromTimestamp    string
//	toTimestamp      string
//	timestamp        string
//	traceID          string
//	spanID           string
//	commit           string
//	parentResourceID string
//)

var rootCmd, regexCmd *cobra.Command
var logger *log.Logger
var flagvalues Models.Flagvalue

func init() {

	rootCmd = &cobra.Command{

		Use:   "dyte",
		Short: "A simple log query processor made for dyte",

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("No args provided")
			}
			if err := flagvalues.CheckDuplicateOnAllFlags(); err != nil {
				logger.Println("duplicate check failed error,", err)
				return err
			}

			runQuery(flagvalues)
			return nil
		},
	}

	regexCmd = &cobra.Command{
		Use:   "regex",
		Short: "Perform regex-based log filtering amde for dyte",

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("No args provided")
			}
			if err := flagvalues.CheckDuplicateOnAllFlags(); err != nil {
				logger.Println("duplicate check failed error,", err)
				return err
			}

			runQuery(flagvalues)
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&flagvalues.Timestamp, "timestamp", "t", "", "Filter logs from  this timestamp")
	rootCmd.Flags().StringVarP(&flagvalues.FromTimestamp, "from", "", "", "Filter logs from timestamp")
	rootCmd.Flags().StringVarP(&flagvalues.ToTimestamp, "to", "", "", "Filter logs to timestamp")

	rootCmd.Flags().StringVarP(&flagvalues.Level.RegularFlag, "level", "l", "", "Filter logs by level")
	rootCmd.Flags().StringVarP(&flagvalues.ResourceId.RegularFlag, "resourceId", "r", "", "Filter logs by resource ID")
	rootCmd.Flags().StringVarP(&flagvalues.Message.RegularFlag, "message", "m", "", "Filter logs by message")
	rootCmd.Flags().StringVarP(&flagvalues.TraceId.RegularFlag, "TraceId", "T", "", "Filter logs by Trace ID")
	rootCmd.Flags().StringVarP(&flagvalues.SpanId.RegularFlag, "spanId", "s", "", "Filter logs by Span ID")
	rootCmd.Flags().StringVarP(&flagvalues.ParentResourceId.RegularFlag, "ParentResId", "p", "", "Filter logs by Parent Resource ID")
	rootCmd.Flags().StringVarP(&flagvalues.Commit.RegularFlag, "commit", "c", "", "Filter logs by commit")

	regexCmd.Flags().StringVarP(&flagvalues.Level.RegexFlag, "level", "l", "", "Filter logs by level pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.ResourceId.RegexFlag, "resourceId", "r", "", "Filter logs by resource ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.Message.RegexFlag, "message", "m", "", "Filter logs by message pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.TraceId.RegexFlag, "TraceId", "T", "", "Filter logs by Trace ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.SpanId.RegexFlag, "spanId", "s", "", "Filter logs by Span ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.ParentResourceId.RegexFlag, "ParentResId", "p", "", "Filter logs by Parent Resource ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.Commit.RegexFlag, "commit", "c", "", "Filter logs by commit pattern using regex")

	rootCmd.AddCommand(regexCmd)
}
func main() {
	logger = comms.LoggerInit()
	rootCmd.TraverseChildren = true //this is required to traverse all the child and parent flags if set to false will only consider child flags
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func runQuery(flagvalue Models.Flagvalue) {
	// Establish a connection to the MySQL database

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
	query, argsList := CLI.PrepareGeneralQuery(flagvalues.Level, flagvalues.Message, flagvalues.ResourceId, flagvalues.TraceId, flagvalues.SpanId, flagvalues.Commit, flagvalues.ParentResourceId, flagvalues.Timestamp)
	fmt.Println(query, argsList)
	logger.Println("creating the prepare statement", query)
	//
	stmt, err := CLI.PrepareGeneralStatement(db, query)
	if err != nil {
		log.Fatal("failed to prepare statements", err)
	}
	logger.Println("Creating prepare statement successful")
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
