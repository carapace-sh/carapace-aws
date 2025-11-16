package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeSubscriptionFiltersCmd = &cobra.Command{
	Use:   "describe-subscription-filters",
	Short: "Lists the subscription filters for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeSubscriptionFiltersCmd).Standalone()

	logs_describeSubscriptionFiltersCmd.Flags().String("filter-name-prefix", "", "The prefix to match.")
	logs_describeSubscriptionFiltersCmd.Flags().String("limit", "", "The maximum number of items returned.")
	logs_describeSubscriptionFiltersCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_describeSubscriptionFiltersCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logs_describeSubscriptionFiltersCmd.MarkFlagRequired("log-group-name")
	logsCmd.AddCommand(logs_describeSubscriptionFiltersCmd)
}
