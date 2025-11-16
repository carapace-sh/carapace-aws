package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listAnalyticsDataAssociationsCmd = &cobra.Command{
	Use:   "list-analytics-data-associations",
	Short: "Lists the association status of requested dataset ID for a given Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listAnalyticsDataAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listAnalyticsDataAssociationsCmd).Standalone()

		connect_listAnalyticsDataAssociationsCmd.Flags().String("data-set-id", "", "The identifier of the dataset to get the association status.")
		connect_listAnalyticsDataAssociationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listAnalyticsDataAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listAnalyticsDataAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listAnalyticsDataAssociationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listAnalyticsDataAssociationsCmd)
}
