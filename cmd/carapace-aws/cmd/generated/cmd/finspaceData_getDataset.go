package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getDatasetCmd = &cobra.Command{
	Use:   "get-dataset",
	Short: "Returns information about a Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getDatasetCmd).Standalone()

		finspaceData_getDatasetCmd.Flags().String("dataset-id", "", "The unique identifier for a Dataset.")
		finspaceData_getDatasetCmd.MarkFlagRequired("dataset-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getDatasetCmd)
}
