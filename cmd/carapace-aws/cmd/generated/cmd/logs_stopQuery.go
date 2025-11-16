package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_stopQueryCmd = &cobra.Command{
	Use:   "stop-query",
	Short: "Stops a CloudWatch Logs Insights query that is in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_stopQueryCmd).Standalone()

	logs_stopQueryCmd.Flags().String("query-id", "", "The ID number of the query to stop.")
	logs_stopQueryCmd.MarkFlagRequired("query-id")
	logsCmd.AddCommand(logs_stopQueryCmd)
}
