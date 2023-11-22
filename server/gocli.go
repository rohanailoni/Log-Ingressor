package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rohanailoni/Log-Ingressor/CLI"
	"github.com/rohanailoni/Log-Ingressor/CLI/Models"
	"github.com/rohanailoni/Log-Ingressor/comms"
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

var rootCmd, regexCmd, authCmd *cobra.Command
var logger *log.Logger
var flagvalues Models.Flagvalue
var FlagUser Models.FlagUser

func init() {

	rootCmd = &cobra.Command{

		Use:   "dyte",
		Short: "A simple log query processor made for dyte",

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("args at root", args)
			if err := CLI.CheckAuthAndPermission(flagvalues); err != nil {
				log.Println("Error while auth")
				return err
			}
			log.Println("Authentication successful")
			if err := flagvalues.CheckDuplicateOnAllFlags(); err != nil {
				logger.Println("duplicate check failed error,", err)
				return err
			}

			runQuery(flagvalues)
			return nil
		},
	}
	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Perform auth for user registered for dyte .If not registerd ask admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if FlagUser.User == "" || FlagUser.Password == "" {

				return errors.New("both username and password are null")
			}
			err := CLI.Authenticate(FlagUser.User, FlagUser.Password)
			if err != nil {
				log.Println("User not authorised for authentications")
				return err
			}
			return nil
		},
	}
	regexCmd = &cobra.Command{
		Use:   "regex",
		Short: "Perform regex-based log filtering made for dyte",

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("args at regex", args)
			if err := CLI.CheckAuthAndPermission(flagvalues); err != nil {
				log.Println("Error while auth")
				return err
			}
			if err := flagvalues.CheckDuplicateOnAllFlags(); err != nil {
				logger.Println("duplicate check failed error,", err)
				return err
			}

			runQuery(flagvalues)
			return nil
		},
	}
	wildCardCmd := &cobra.Command{
		Use:   "wildcard",
		Short: "Perform wilcardbased log filtering made for dyte",

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("args at wildcard", args)
			if err := CLI.CheckAuthAndPermission(flagvalues); err != nil {
				log.Println("Error while auth")
				return err
			}
			if err := flagvalues.CheckDuplicateOnAllFlags(); err != nil {
				logger.Println("duplicate check failed error,", err)
				return err
			}

			runQuery(flagvalues)
			return nil
		},
	}
	rootCmd.Flags().StringVarP(&flagvalues.Level.RegularFlag, "level", "l", "", "Filter logs by level")
	rootCmd.Flags().StringVarP(&flagvalues.ResourceId.RegularFlag, "resourceId", "r", "", "Filter logs by resource ID")
	rootCmd.Flags().StringVarP(&flagvalues.Message.RegularFlag, "message", "m", "", "Filter logs by message")
	rootCmd.Flags().StringVarP(&flagvalues.TraceId.RegularFlag, "TraceId", "T", "", "Filter logs by Trace ID")
	rootCmd.Flags().StringVarP(&flagvalues.SpanId.RegularFlag, "spanId", "s", "", "Filter logs by Span ID")
	rootCmd.Flags().StringVarP(&flagvalues.ParentResourceId.RegularFlag, "ParentResId", "p", "", "Filter logs by Parent Resource ID")
	rootCmd.Flags().StringVarP(&flagvalues.Commit.RegularFlag, "commit", "c", "", "Filter logs by commit")
	//time flags
	rootCmd.Flags().StringVarP(&flagvalues.Timestamp, "timestamp", "t", "", "Filter logs from  this timestamp")
	rootCmd.Flags().StringVar(&flagvalues.FromTimestamp, "from", "", "Filter logs from timestamp")
	rootCmd.Flags().StringVar(&flagvalues.ToTimestamp, "to", "", "Filter logs to timestamp")

	//auth flags
	authCmd.Flags().StringVarP(&FlagUser.User, "user", "u", "", "Username for auth")
	authCmd.Flags().StringVarP(&FlagUser.Password, "password", "p", "", "Password for auth")

	regexCmd.Flags().StringVarP(&flagvalues.Level.RegexFlag, "level", "l", "", "Filter logs by level pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.ResourceId.RegexFlag, "resourceId", "r", "", "Filter logs by resource ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.Message.RegexFlag, "message", "m", "", "Filter logs by message pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.TraceId.RegexFlag, "TraceId", "T", "", "Filter logs by Trace ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.SpanId.RegexFlag, "spanId", "s", "", "Filter logs by Span ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.ParentResourceId.RegexFlag, "ParentResId", "p", "", "Filter logs by Parent Resource ID pattern using regex")
	regexCmd.Flags().StringVarP(&flagvalues.Commit.RegexFlag, "commit", "c", "", "Filter logs by commit pattern using regex")

	wildCardCmd.Flags().StringVarP(&flagvalues.Level.WildcardFlag, "level", "l", "", "Filter logs by level pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.ResourceId.WildcardFlag, "resourceId", "r", "", "Filter logs by resource ID pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.Message.WildcardFlag, "message", "m", "", "Filter logs by message pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.TraceId.WildcardFlag, "TraceId", "T", "", "Filter logs by Trace ID pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.SpanId.WildcardFlag, "spanId", "s", "", "Filter logs by Span ID pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.ParentResourceId.WildcardFlag, "ParentResId", "p", "", "Filter logs by Parent Resource ID pattern using wildcard(sql LIKE)")
	wildCardCmd.Flags().StringVarP(&flagvalues.Commit.WildcardFlag, "commit", "c", "", "Filter logs by commit pattern using wildcard(sql LIKE)")

	rootCmd.AddCommand(regexCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(wildCardCmd)
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
	query, argsList := CLI.PrepareGeneralQuery(flagvalues.Level, flagvalues.Message, flagvalues.ResourceId, flagvalues.TraceId, flagvalues.SpanId, flagvalues.Commit, flagvalues.ParentResourceId, flagvalues.Timestamp, flagvalue.FromTimestamp, flagvalue.ToTimestamp)
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
