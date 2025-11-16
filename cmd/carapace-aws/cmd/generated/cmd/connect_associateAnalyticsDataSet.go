package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateAnalyticsDataSetCmd = &cobra.Command{
	Use:   "associate-analytics-data-set",
	Short: "Associates the specified dataset for a Amazon Connect instance with the target account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateAnalyticsDataSetCmd).Standalone()

	connect_associateAnalyticsDataSetCmd.Flags().String("data-set-id", "", "The identifier of the dataset to associate with the target account.")
	connect_associateAnalyticsDataSetCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateAnalyticsDataSetCmd.Flags().String("target-account-id", "", "The identifier of the target account.")
	connect_associateAnalyticsDataSetCmd.MarkFlagRequired("data-set-id")
	connect_associateAnalyticsDataSetCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_associateAnalyticsDataSetCmd)
}
