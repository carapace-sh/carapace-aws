package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listChangesetsCmd = &cobra.Command{
	Use:   "list-changesets",
	Short: "Lists the FinSpace Changesets for a Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listChangesetsCmd).Standalone()

	finspaceData_listChangesetsCmd.Flags().String("dataset-id", "", "The unique identifier for the FinSpace Dataset to which the Changeset belongs.")
	finspaceData_listChangesetsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listChangesetsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listChangesetsCmd.MarkFlagRequired("dataset-id")
	finspaceDataCmd.AddCommand(finspaceData_listChangesetsCmd)
}
