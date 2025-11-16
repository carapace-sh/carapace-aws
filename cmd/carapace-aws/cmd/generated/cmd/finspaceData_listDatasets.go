package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Lists all of the active Datasets that a user has access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_listDatasetsCmd).Standalone()

		finspaceData_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
		finspaceData_listDatasetsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	})
	finspaceDataCmd.AddCommand(finspaceData_listDatasetsCmd)
}
