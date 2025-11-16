package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_startQueryCmd = &cobra.Command{
	Use:   "start-query",
	Short: "Starts a query of one or more log groups using CloudWatch Logs Insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_startQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_startQueryCmd).Standalone()

		logs_startQueryCmd.Flags().String("end-time", "", "The end of the time range to query.")
		logs_startQueryCmd.Flags().String("limit", "", "The maximum number of log events to return in the query.")
		logs_startQueryCmd.Flags().String("log-group-identifiers", "", "The list of log groups to query.")
		logs_startQueryCmd.Flags().String("log-group-name", "", "The log group on which to perform the query.")
		logs_startQueryCmd.Flags().String("log-group-names", "", "The list of log groups to be queried.")
		logs_startQueryCmd.Flags().String("query-language", "", "Specify the query language to use for this query.")
		logs_startQueryCmd.Flags().String("query-string", "", "The query string to use.")
		logs_startQueryCmd.Flags().String("start-time", "", "The beginning of the time range to query.")
		logs_startQueryCmd.MarkFlagRequired("end-time")
		logs_startQueryCmd.MarkFlagRequired("query-string")
		logs_startQueryCmd.MarkFlagRequired("start-time")
	})
	logsCmd.AddCommand(logs_startQueryCmd)
}
