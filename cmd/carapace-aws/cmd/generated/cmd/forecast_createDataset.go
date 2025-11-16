package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates an Amazon Forecast dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createDatasetCmd).Standalone()

	forecast_createDatasetCmd.Flags().String("data-frequency", "", "The frequency of data collection.")
	forecast_createDatasetCmd.Flags().String("dataset-name", "", "A name for the dataset.")
	forecast_createDatasetCmd.Flags().String("dataset-type", "", "The dataset type.")
	forecast_createDatasetCmd.Flags().String("domain", "", "The domain associated with the dataset.")
	forecast_createDatasetCmd.Flags().String("encryption-config", "", "An Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.")
	forecast_createDatasetCmd.Flags().String("schema", "", "The schema for the dataset.")
	forecast_createDatasetCmd.Flags().String("tags", "", "The optional metadata that you apply to the dataset to help you categorize and organize them.")
	forecast_createDatasetCmd.MarkFlagRequired("dataset-name")
	forecast_createDatasetCmd.MarkFlagRequired("dataset-type")
	forecast_createDatasetCmd.MarkFlagRequired("domain")
	forecast_createDatasetCmd.MarkFlagRequired("schema")
	forecastCmd.AddCommand(forecast_createDatasetCmd)
}
