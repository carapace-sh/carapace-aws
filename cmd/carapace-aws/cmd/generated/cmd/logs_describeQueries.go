package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeQueriesCmd = &cobra.Command{
	Use:   "describe-queries",
	Short: "Returns a list of CloudWatch Logs Insights queries that are scheduled, running, or have been run recently in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeQueriesCmd).Standalone()

	logs_describeQueriesCmd.Flags().String("log-group-name", "", "Limits the returned queries to only those for the specified log group.")
	logs_describeQueriesCmd.Flags().String("max-results", "", "Limits the number of returned queries to the specified number.")
	logs_describeQueriesCmd.Flags().String("next-token", "", "")
	logs_describeQueriesCmd.Flags().String("query-language", "", "Limits the returned queries to only the queries that use the specified query language.")
	logs_describeQueriesCmd.Flags().String("status", "", "Limits the returned queries to only those that have the specified status.")
	logsCmd.AddCommand(logs_describeQueriesCmd)
}
