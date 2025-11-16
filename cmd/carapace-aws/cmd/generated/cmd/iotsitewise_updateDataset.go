package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateDatasetCmd = &cobra.Command{
	Use:   "update-dataset",
	Short: "Updates a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateDatasetCmd).Standalone()

	iotsitewise_updateDatasetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateDatasetCmd.Flags().String("dataset-description", "", "A description about the dataset, and its functionality.")
	iotsitewise_updateDatasetCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
	iotsitewise_updateDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset.")
	iotsitewise_updateDatasetCmd.Flags().String("dataset-source", "", "The data source for the dataset.")
	iotsitewise_updateDatasetCmd.MarkFlagRequired("dataset-id")
	iotsitewise_updateDatasetCmd.MarkFlagRequired("dataset-name")
	iotsitewise_updateDatasetCmd.MarkFlagRequired("dataset-source")
	iotsitewiseCmd.AddCommand(iotsitewise_updateDatasetCmd)
}
