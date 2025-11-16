package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_updateDatasetGroupCmd = &cobra.Command{
	Use:   "update-dataset-group",
	Short: "Replaces the datasets in a dataset group with the specified datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_updateDatasetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_updateDatasetGroupCmd).Standalone()

		forecast_updateDatasetGroupCmd.Flags().String("dataset-arns", "", "An array of the Amazon Resource Names (ARNs) of the datasets to add to the dataset group.")
		forecast_updateDatasetGroupCmd.Flags().String("dataset-group-arn", "", "The ARN of the dataset group.")
		forecast_updateDatasetGroupCmd.MarkFlagRequired("dataset-arns")
		forecast_updateDatasetGroupCmd.MarkFlagRequired("dataset-group-arn")
	})
	forecastCmd.AddCommand(forecast_updateDatasetGroupCmd)
}
