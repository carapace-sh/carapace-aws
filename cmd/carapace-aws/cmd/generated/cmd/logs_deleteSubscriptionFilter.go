package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteSubscriptionFilterCmd = &cobra.Command{
	Use:   "delete-subscription-filter",
	Short: "Deletes the specified subscription filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteSubscriptionFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteSubscriptionFilterCmd).Standalone()

		logs_deleteSubscriptionFilterCmd.Flags().String("filter-name", "", "The name of the subscription filter.")
		logs_deleteSubscriptionFilterCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_deleteSubscriptionFilterCmd.MarkFlagRequired("filter-name")
		logs_deleteSubscriptionFilterCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_deleteSubscriptionFilterCmd)
}
