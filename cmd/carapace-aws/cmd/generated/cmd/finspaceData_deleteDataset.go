package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes a FinSpace Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_deleteDatasetCmd).Standalone()

	finspaceData_deleteDatasetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_deleteDatasetCmd.Flags().String("dataset-id", "", "The unique identifier of the Dataset to be deleted.")
	finspaceData_deleteDatasetCmd.MarkFlagRequired("dataset-id")
	finspaceDataCmd.AddCommand(finspaceData_deleteDatasetCmd)
}
