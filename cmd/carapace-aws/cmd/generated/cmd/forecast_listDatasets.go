package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Returns a list of datasets created using the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listDatasetsCmd).Standalone()

		forecast_listDatasetsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listDatasetsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	})
	forecastCmd.AddCommand(forecast_listDatasetsCmd)
}
