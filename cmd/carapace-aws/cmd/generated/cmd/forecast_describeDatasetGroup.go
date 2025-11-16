package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeDatasetGroupCmd = &cobra.Command{
	Use:   "describe-dataset-group",
	Short: "Describes a dataset group created using the [CreateDatasetGroup](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeDatasetGroupCmd).Standalone()

	forecast_describeDatasetGroupCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group.")
	forecast_describeDatasetGroupCmd.MarkFlagRequired("dataset-group-arn")
	forecastCmd.AddCommand(forecast_describeDatasetGroupCmd)
}
