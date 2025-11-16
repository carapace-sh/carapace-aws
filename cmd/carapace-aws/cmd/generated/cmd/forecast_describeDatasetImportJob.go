package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeDatasetImportJobCmd = &cobra.Command{
	Use:   "describe-dataset-import-job",
	Short: "Describes a dataset import job created using the [CreateDatasetImportJob](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeDatasetImportJobCmd).Standalone()

	forecast_describeDatasetImportJobCmd.Flags().String("dataset-import-job-arn", "", "The Amazon Resource Name (ARN) of the dataset import job.")
	forecast_describeDatasetImportJobCmd.MarkFlagRequired("dataset-import-job-arn")
	forecastCmd.AddCommand(forecast_describeDatasetImportJobCmd)
}
