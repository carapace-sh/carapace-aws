package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_deleteDatasetContentCmd = &cobra.Command{
	Use:   "delete-dataset-content",
	Short: "Deletes the content of the specified dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_deleteDatasetContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_deleteDatasetContentCmd).Standalone()

		iotanalytics_deleteDatasetContentCmd.Flags().String("dataset-name", "", "The name of the dataset whose content is deleted.")
		iotanalytics_deleteDatasetContentCmd.Flags().String("version-id", "", "The version of the dataset whose content is deleted.")
		iotanalytics_deleteDatasetContentCmd.MarkFlagRequired("dataset-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_deleteDatasetContentCmd)
}
