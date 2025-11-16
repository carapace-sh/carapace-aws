package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listDatasetGroupsCmd = &cobra.Command{
	Use:   "list-dataset-groups",
	Short: "Returns a list of dataset groups created using the [CreateDatasetGroup](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listDatasetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listDatasetGroupsCmd).Standalone()

		forecast_listDatasetGroupsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listDatasetGroupsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	})
	forecastCmd.AddCommand(forecast_listDatasetGroupsCmd)
}
