package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes the specified dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_deleteDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_deleteDatasetCmd).Standalone()

		iotanalytics_deleteDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset to delete.")
		iotanalytics_deleteDatasetCmd.MarkFlagRequired("dataset-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_deleteDatasetCmd)
}
