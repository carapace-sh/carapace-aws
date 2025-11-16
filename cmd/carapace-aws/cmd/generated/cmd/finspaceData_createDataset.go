package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates a new FinSpace Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_createDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_createDatasetCmd).Standalone()

		finspaceData_createDatasetCmd.Flags().String("alias", "", "The unique resource identifier for a Dataset.")
		finspaceData_createDatasetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_createDatasetCmd.Flags().String("dataset-description", "", "Description of a Dataset.")
		finspaceData_createDatasetCmd.Flags().String("dataset-title", "", "Display title for a FinSpace Dataset.")
		finspaceData_createDatasetCmd.Flags().String("kind", "", "The format in which Dataset data is structured.")
		finspaceData_createDatasetCmd.Flags().String("owner-info", "", "Contact information for a Dataset owner.")
		finspaceData_createDatasetCmd.Flags().String("permission-group-params", "", "Permission group parameters for Dataset permissions.")
		finspaceData_createDatasetCmd.Flags().String("schema-definition", "", "Definition for a schema on a tabular Dataset.")
		finspaceData_createDatasetCmd.MarkFlagRequired("dataset-title")
		finspaceData_createDatasetCmd.MarkFlagRequired("kind")
		finspaceData_createDatasetCmd.MarkFlagRequired("permission-group-params")
	})
	finspaceDataCmd.AddCommand(finspaceData_createDatasetCmd)
}
