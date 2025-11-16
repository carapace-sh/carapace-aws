package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getChangesetCmd = &cobra.Command{
	Use:   "get-changeset",
	Short: "Get information about a Changeset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getChangesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getChangesetCmd).Standalone()

		finspaceData_getChangesetCmd.Flags().String("changeset-id", "", "The unique identifier of the Changeset for which to get data.")
		finspaceData_getChangesetCmd.Flags().String("dataset-id", "", "The unique identifier for the FinSpace Dataset where the Changeset is created.")
		finspaceData_getChangesetCmd.MarkFlagRequired("changeset-id")
		finspaceData_getChangesetCmd.MarkFlagRequired("dataset-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getChangesetCmd)
}
