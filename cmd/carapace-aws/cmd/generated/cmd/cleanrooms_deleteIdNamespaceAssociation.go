package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteIdNamespaceAssociationCmd = &cobra.Command{
	Use:   "delete-id-namespace-association",
	Short: "Deletes an ID namespace association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteIdNamespaceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteIdNamespaceAssociationCmd).Standalone()

		cleanrooms_deleteIdNamespaceAssociationCmd.Flags().String("id-namespace-association-identifier", "", "The unique identifier of the ID namespace association that you want to delete.")
		cleanrooms_deleteIdNamespaceAssociationCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID namespace association that you want to delete.")
		cleanrooms_deleteIdNamespaceAssociationCmd.MarkFlagRequired("id-namespace-association-identifier")
		cleanrooms_deleteIdNamespaceAssociationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteIdNamespaceAssociationCmd)
}
