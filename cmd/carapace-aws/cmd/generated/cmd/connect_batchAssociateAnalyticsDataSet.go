package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_batchAssociateAnalyticsDataSetCmd = &cobra.Command{
	Use:   "batch-associate-analytics-data-set",
	Short: "Associates a list of analytics datasets for a given Amazon Connect instance to a target account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_batchAssociateAnalyticsDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_batchAssociateAnalyticsDataSetCmd).Standalone()

		connect_batchAssociateAnalyticsDataSetCmd.Flags().String("data-set-ids", "", "An array of dataset identifiers to associate.")
		connect_batchAssociateAnalyticsDataSetCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_batchAssociateAnalyticsDataSetCmd.Flags().String("target-account-id", "", "The identifier of the target account.")
		connect_batchAssociateAnalyticsDataSetCmd.MarkFlagRequired("data-set-ids")
		connect_batchAssociateAnalyticsDataSetCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_batchAssociateAnalyticsDataSetCmd)
}
