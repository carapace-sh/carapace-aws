package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createDatasetImportJobCmd = &cobra.Command{
	Use:   "create-dataset-import-job",
	Short: "Imports your training data to an Amazon Forecast dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createDatasetImportJobCmd).Standalone()

	forecast_createDatasetImportJobCmd.Flags().String("data-source", "", "The location of the training data to import and an Identity and Access Management (IAM) role that Amazon Forecast can assume to access the data.")
	forecast_createDatasetImportJobCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the Amazon Forecast dataset that you want to import data to.")
	forecast_createDatasetImportJobCmd.Flags().String("dataset-import-job-name", "", "The name for the dataset import job.")
	forecast_createDatasetImportJobCmd.Flags().String("format", "", "The format of the imported data, CSV or PARQUET.")
	forecast_createDatasetImportJobCmd.Flags().String("geolocation-format", "", "The format of the geolocation attribute.")
	forecast_createDatasetImportJobCmd.Flags().String("import-mode", "", "Specifies whether the dataset import job is a `FULL` or `INCREMENTAL` import.")
	forecast_createDatasetImportJobCmd.Flags().String("tags", "", "The optional metadata that you apply to the dataset import job to help you categorize and organize them.")
	forecast_createDatasetImportJobCmd.Flags().String("time-zone", "", "A single time zone for every item in your dataset.")
	forecast_createDatasetImportJobCmd.Flags().String("timestamp-format", "", "The format of timestamps in the dataset.")
	forecast_createDatasetImportJobCmd.Flags().String("use-geolocation-for-time-zone", "", "Automatically derive time zone information from the geolocation attribute.")
	forecast_createDatasetImportJobCmd.MarkFlagRequired("data-source")
	forecast_createDatasetImportJobCmd.MarkFlagRequired("dataset-arn")
	forecast_createDatasetImportJobCmd.MarkFlagRequired("dataset-import-job-name")
	forecastCmd.AddCommand(forecast_createDatasetImportJobCmd)
}
