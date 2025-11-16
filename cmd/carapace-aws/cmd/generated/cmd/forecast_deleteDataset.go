package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes an Amazon Forecast dataset that was created using the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deleteDatasetCmd).Standalone()

		forecast_deleteDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to delete.")
		forecast_deleteDatasetCmd.MarkFlagRequired("dataset-arn")
	})
	forecastCmd.AddCommand(forecast_deleteDatasetCmd)
}
