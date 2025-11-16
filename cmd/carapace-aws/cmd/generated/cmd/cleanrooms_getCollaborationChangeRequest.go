package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationChangeRequestCmd = &cobra.Command{
	Use:   "get-collaboration-change-request",
	Short: "Retrieves detailed information about a specific collaboration change request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationChangeRequestCmd).Standalone()

	cleanrooms_getCollaborationChangeRequestCmd.Flags().String("change-request-identifier", "", "A unique identifier for the change request to retrieve.")
	cleanrooms_getCollaborationChangeRequestCmd.Flags().String("collaboration-identifier", "", "The identifier of the collaboration that the change request is made against.")
	cleanrooms_getCollaborationChangeRequestCmd.MarkFlagRequired("change-request-identifier")
	cleanrooms_getCollaborationChangeRequestCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationChangeRequestCmd)
}
