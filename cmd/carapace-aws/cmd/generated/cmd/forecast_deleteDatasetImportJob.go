package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteDatasetImportJobCmd = &cobra.Command{
	Use:   "delete-dataset-import-job",
	Short: "Deletes a dataset import job created using the [CreateDatasetImportJob](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteDatasetImportJobCmd).Standalone()

	forecast_deleteDatasetImportJobCmd.Flags().String("dataset-import-job-arn", "", "The Amazon Resource Name (ARN) of the dataset import job to delete.")
	forecast_deleteDatasetImportJobCmd.MarkFlagRequired("dataset-import-job-arn")
	forecastCmd.AddCommand(forecast_deleteDatasetImportJobCmd)
}
