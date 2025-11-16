package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_executeScheduledQueryCmd = &cobra.Command{
	Use:   "execute-scheduled-query",
	Short: "You can use this API to run a scheduled query manually.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_executeScheduledQueryCmd).Standalone()

	timestreamQuery_executeScheduledQueryCmd.Flags().String("client-token", "", "Not used.")
	timestreamQuery_executeScheduledQueryCmd.Flags().String("invocation-time", "", "The timestamp in UTC.")
	timestreamQuery_executeScheduledQueryCmd.Flags().String("query-insights", "", "Encapsulates settings for enabling `QueryInsights`.")
	timestreamQuery_executeScheduledQueryCmd.Flags().String("scheduled-query-arn", "", "ARN of the scheduled query.")
	timestreamQuery_executeScheduledQueryCmd.MarkFlagRequired("invocation-time")
	timestreamQuery_executeScheduledQueryCmd.MarkFlagRequired("scheduled-query-arn")
	timestreamQueryCmd.AddCommand(timestreamQuery_executeScheduledQueryCmd)
}
