package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes a dataset and associated artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteDatasetCmd).Standalone()

	lookoutequipment_deleteDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset to be deleted.")
	lookoutequipment_deleteDatasetCmd.MarkFlagRequired("dataset-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteDatasetCmd)
}
