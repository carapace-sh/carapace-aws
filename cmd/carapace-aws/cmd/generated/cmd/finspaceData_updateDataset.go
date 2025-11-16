package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_updateDatasetCmd = &cobra.Command{
	Use:   "update-dataset",
	Short: "Updates a FinSpace Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_updateDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_updateDatasetCmd).Standalone()

		finspaceData_updateDatasetCmd.Flags().String("alias", "", "The unique resource identifier for a Dataset.")
		finspaceData_updateDatasetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_updateDatasetCmd.Flags().String("dataset-description", "", "A description for the Dataset.")
		finspaceData_updateDatasetCmd.Flags().String("dataset-id", "", "The unique identifier for the Dataset to update.")
		finspaceData_updateDatasetCmd.Flags().String("dataset-title", "", "A display title for the Dataset.")
		finspaceData_updateDatasetCmd.Flags().String("kind", "", "The format in which the Dataset data is structured.")
		finspaceData_updateDatasetCmd.Flags().String("schema-definition", "", "Definition for a schema on a tabular Dataset.")
		finspaceData_updateDatasetCmd.MarkFlagRequired("dataset-id")
		finspaceData_updateDatasetCmd.MarkFlagRequired("dataset-title")
		finspaceData_updateDatasetCmd.MarkFlagRequired("kind")
	})
	finspaceDataCmd.AddCommand(finspaceData_updateDatasetCmd)
}
