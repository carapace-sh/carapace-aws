package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteDatasetCmd).Standalone()

		iotsitewise_deleteDatasetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteDatasetCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
		iotsitewise_deleteDatasetCmd.MarkFlagRequired("dataset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteDatasetCmd)
}
