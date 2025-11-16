package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_getDatasetContentCmd = &cobra.Command{
	Use:   "get-dataset-content",
	Short: "Retrieves the contents of a dataset as presigned URIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_getDatasetContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_getDatasetContentCmd).Standalone()

		iotanalytics_getDatasetContentCmd.Flags().String("dataset-name", "", "The name of the dataset whose contents are retrieved.")
		iotanalytics_getDatasetContentCmd.Flags().String("version-id", "", "The version of the dataset whose contents are retrieved.")
		iotanalytics_getDatasetContentCmd.MarkFlagRequired("dataset-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_getDatasetContentCmd)
}
