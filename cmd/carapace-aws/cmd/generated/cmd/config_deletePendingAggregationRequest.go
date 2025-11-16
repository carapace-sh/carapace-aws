package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deletePendingAggregationRequestCmd = &cobra.Command{
	Use:   "delete-pending-aggregation-request",
	Short: "Deletes pending authorization requests for a specified aggregator account in a specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deletePendingAggregationRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deletePendingAggregationRequestCmd).Standalone()

		config_deletePendingAggregationRequestCmd.Flags().String("requester-account-id", "", "The 12-digit account ID of the account requesting to aggregate data.")
		config_deletePendingAggregationRequestCmd.Flags().String("requester-aws-region", "", "The region requesting to aggregate data.")
		config_deletePendingAggregationRequestCmd.MarkFlagRequired("requester-account-id")
		config_deletePendingAggregationRequestCmd.MarkFlagRequired("requester-aws-region")
	})
	configCmd.AddCommand(config_deletePendingAggregationRequestCmd)
}
