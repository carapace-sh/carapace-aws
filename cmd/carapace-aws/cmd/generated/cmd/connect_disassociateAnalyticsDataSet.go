package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateAnalyticsDataSetCmd = &cobra.Command{
	Use:   "disassociate-analytics-data-set",
	Short: "Removes the dataset ID associated with a given Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateAnalyticsDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateAnalyticsDataSetCmd).Standalone()

		connect_disassociateAnalyticsDataSetCmd.Flags().String("data-set-id", "", "The identifier of the dataset to remove.")
		connect_disassociateAnalyticsDataSetCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateAnalyticsDataSetCmd.Flags().String("target-account-id", "", "The identifier of the target account.")
		connect_disassociateAnalyticsDataSetCmd.MarkFlagRequired("data-set-id")
		connect_disassociateAnalyticsDataSetCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_disassociateAnalyticsDataSetCmd)
}
