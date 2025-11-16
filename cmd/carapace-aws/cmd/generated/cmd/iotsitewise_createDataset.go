package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates a dataset to connect an external datasource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createDatasetCmd).Standalone()

	iotsitewise_createDatasetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createDatasetCmd.Flags().String("dataset-description", "", "A description about the dataset, and its functionality.")
	iotsitewise_createDatasetCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
	iotsitewise_createDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset.")
	iotsitewise_createDatasetCmd.Flags().String("dataset-source", "", "The data source for the dataset.")
	iotsitewise_createDatasetCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the access policy.")
	iotsitewise_createDatasetCmd.MarkFlagRequired("dataset-name")
	iotsitewise_createDatasetCmd.MarkFlagRequired("dataset-source")
	iotsitewiseCmd.AddCommand(iotsitewise_createDatasetCmd)
}
