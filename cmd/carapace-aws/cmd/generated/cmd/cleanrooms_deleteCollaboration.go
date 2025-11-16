package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteCollaborationCmd = &cobra.Command{
	Use:   "delete-collaboration",
	Short: "Deletes a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteCollaborationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteCollaborationCmd).Standalone()

		cleanrooms_deleteCollaborationCmd.Flags().String("collaboration-identifier", "", "The identifier for the collaboration.")
		cleanrooms_deleteCollaborationCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteCollaborationCmd)
}
