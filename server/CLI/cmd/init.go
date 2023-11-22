package cmd

import (
	"fmt"
	"github.com/rohanailoni/Log-Ingressor/CLI/Models"
	"github.com/spf13/cobra"
)

//var (
//	level            string
//	messagePattern   string
//	commitPattern    string
//	resourceID       string
//	timestamp        string
//	fromTimestamp    string
//	toTimestamp      string
//	traceID          string
//	spanID           string
//	parentResourceID string
//	message          string
//)

func InitCli() (*cobra.Command, *cobra.Command, Models.Flagvalue, error) {
	var flagvalues Models.Flagvalue
	rootCmd := &cobra.Command{

		Use:   "dyte",
		Short: "A simple log query processor made for dyte",
		Run: func(cmd *cobra.Command, args []string) {
			flagvalues.Timestamp, _ = cmd.Flags().GetString("timestamp")
			flagvalues.FromTimestamp, _ = cmd.Flags().GetString("from")
			flagvalues.Level.RegexFlag, _ = cmd.Flags().GetString("level")
			fmt.Println(flagvalues.Level.RegexFlag, cmd.Flags())
		},
	}
	fmt.Println("we are in inti", flagvalues.Level.RegexFlag)
	regexCmd := &cobra.Command{
		Use:   "regex",
		Short: "Perform regex-based log filtering amde for dyte",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
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
	fmt.Println("we are in inti", flagvalues)
	return rootCmd, regexCmd, flagvalues, flagvalues.CheckDuplicateOnAllFlags()

}
