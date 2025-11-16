package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listDatasetImportJobsCmd = &cobra.Command{
	Use:   "list-dataset-import-jobs",
	Short: "Returns a list of dataset import jobs created using the [CreateDatasetImportJob](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listDatasetImportJobsCmd).Standalone()

	forecast_listDatasetImportJobsCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listDatasetImportJobsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
	forecast_listDatasetImportJobsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastCmd.AddCommand(forecast_listDatasetImportJobsCmd)
}
