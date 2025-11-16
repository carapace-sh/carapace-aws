package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createDatasetGroupCmd = &cobra.Command{
	Use:   "create-dataset-group",
	Short: "Creates a dataset group, which holds a collection of related datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createDatasetGroupCmd).Standalone()

	forecast_createDatasetGroupCmd.Flags().String("dataset-arns", "", "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.")
	forecast_createDatasetGroupCmd.Flags().String("dataset-group-name", "", "A name for the dataset group.")
	forecast_createDatasetGroupCmd.Flags().String("domain", "", "The domain associated with the dataset group.")
	forecast_createDatasetGroupCmd.Flags().String("tags", "", "The optional metadata that you apply to the dataset group to help you categorize and organize them.")
	forecast_createDatasetGroupCmd.MarkFlagRequired("dataset-group-name")
	forecast_createDatasetGroupCmd.MarkFlagRequired("domain")
	forecastCmd.AddCommand(forecast_createDatasetGroupCmd)
}
