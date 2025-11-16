package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates a dataset to upload training or test data for a model associated with a flywheel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_createDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_createDatasetCmd).Standalone()

		comprehend_createDatasetCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_createDatasetCmd.Flags().String("dataset-name", "", "Name of the dataset.")
		comprehend_createDatasetCmd.Flags().String("dataset-type", "", "The dataset type.")
		comprehend_createDatasetCmd.Flags().String("description", "", "Description of the dataset.")
		comprehend_createDatasetCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel of the flywheel to receive the data.")
		comprehend_createDatasetCmd.Flags().String("input-data-config", "", "Information about the input data configuration.")
		comprehend_createDatasetCmd.Flags().String("tags", "", "Tags for the dataset.")
		comprehend_createDatasetCmd.MarkFlagRequired("dataset-name")
		comprehend_createDatasetCmd.MarkFlagRequired("flywheel-arn")
		comprehend_createDatasetCmd.MarkFlagRequired("input-data-config")
	})
	comprehendCmd.AddCommand(comprehend_createDatasetCmd)
}
