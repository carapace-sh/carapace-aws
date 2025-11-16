package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_createDatasetContentCmd = &cobra.Command{
	Use:   "create-dataset-content",
	Short: "Creates the content of a dataset by applying a `queryAction` (a SQL query) or a `containerAction` (executing a containerized application).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_createDatasetContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_createDatasetContentCmd).Standalone()

		iotanalytics_createDatasetContentCmd.Flags().String("dataset-name", "", "The name of the dataset.")
		iotanalytics_createDatasetContentCmd.Flags().String("version-id", "", "The version ID of the dataset content.")
		iotanalytics_createDatasetContentCmd.MarkFlagRequired("dataset-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_createDatasetContentCmd)
}
