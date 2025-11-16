package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_batchDisassociateAnalyticsDataSetCmd = &cobra.Command{
	Use:   "batch-disassociate-analytics-data-set",
	Short: "Removes a list of analytics datasets associated with a given Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_batchDisassociateAnalyticsDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_batchDisassociateAnalyticsDataSetCmd).Standalone()

		connect_batchDisassociateAnalyticsDataSetCmd.Flags().String("data-set-ids", "", "An array of associated dataset identifiers to remove.")
		connect_batchDisassociateAnalyticsDataSetCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_batchDisassociateAnalyticsDataSetCmd.Flags().String("target-account-id", "", "The identifier of the target account.")
		connect_batchDisassociateAnalyticsDataSetCmd.MarkFlagRequired("data-set-ids")
		connect_batchDisassociateAnalyticsDataSetCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_batchDisassociateAnalyticsDataSetCmd)
}
