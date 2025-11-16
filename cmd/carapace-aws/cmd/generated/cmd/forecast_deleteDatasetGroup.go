package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteDatasetGroupCmd = &cobra.Command{
	Use:   "delete-dataset-group",
	Short: "Deletes a dataset group created using the [CreateDatasetGroup](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html) operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteDatasetGroupCmd).Standalone()

	forecast_deleteDatasetGroupCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group to delete.")
	forecast_deleteDatasetGroupCmd.MarkFlagRequired("dataset-group-arn")
	forecastCmd.AddCommand(forecast_deleteDatasetGroupCmd)
}
