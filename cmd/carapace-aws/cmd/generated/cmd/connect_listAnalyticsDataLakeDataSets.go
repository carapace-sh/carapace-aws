package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listAnalyticsDataLakeDataSetsCmd = &cobra.Command{
	Use:   "list-analytics-data-lake-data-sets",
	Short: "Lists the data lake datasets available to associate with for a given Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listAnalyticsDataLakeDataSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listAnalyticsDataLakeDataSetsCmd).Standalone()

		connect_listAnalyticsDataLakeDataSetsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listAnalyticsDataLakeDataSetsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listAnalyticsDataLakeDataSetsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listAnalyticsDataLakeDataSetsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listAnalyticsDataLakeDataSetsCmd)
}
