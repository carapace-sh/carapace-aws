package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationIdNamespaceAssociationCmd = &cobra.Command{
	Use:   "get-collaboration-id-namespace-association",
	Short: "Retrieves an ID namespace association from a specific collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationIdNamespaceAssociationCmd).Standalone()

	cleanrooms_getCollaborationIdNamespaceAssociationCmd.Flags().String("collaboration-identifier", "", "The unique identifier of the collaboration that contains the ID namespace association that you want to retrieve.")
	cleanrooms_getCollaborationIdNamespaceAssociationCmd.Flags().String("id-namespace-association-identifier", "", "The unique identifier of the ID namespace association that you want to retrieve.")
	cleanrooms_getCollaborationIdNamespaceAssociationCmd.MarkFlagRequired("collaboration-identifier")
	cleanrooms_getCollaborationIdNamespaceAssociationCmd.MarkFlagRequired("id-namespace-association-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationIdNamespaceAssociationCmd)
}
