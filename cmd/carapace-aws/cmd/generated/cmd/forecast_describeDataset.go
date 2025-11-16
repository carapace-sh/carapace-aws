package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Describes an Amazon Forecast dataset created using the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeDatasetCmd).Standalone()

		forecast_describeDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset.")
		forecast_describeDatasetCmd.MarkFlagRequired("dataset-arn")
	})
	forecastCmd.AddCommand(forecast_describeDatasetCmd)
}
