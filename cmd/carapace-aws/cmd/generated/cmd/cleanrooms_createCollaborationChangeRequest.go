package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createCollaborationChangeRequestCmd = &cobra.Command{
	Use:   "create-collaboration-change-request",
	Short: "Creates a new change request to modify an existing collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createCollaborationChangeRequestCmd).Standalone()

	cleanrooms_createCollaborationChangeRequestCmd.Flags().String("changes", "", "The list of changes to apply to the collaboration.")
	cleanrooms_createCollaborationChangeRequestCmd.Flags().String("collaboration-identifier", "", "The identifier of the collaboration that the change request is made against.")
	cleanrooms_createCollaborationChangeRequestCmd.MarkFlagRequired("changes")
	cleanrooms_createCollaborationChangeRequestCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_createCollaborationChangeRequestCmd)
}
