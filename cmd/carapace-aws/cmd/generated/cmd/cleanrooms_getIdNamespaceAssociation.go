package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getIdNamespaceAssociationCmd = &cobra.Command{
	Use:   "get-id-namespace-association",
	Short: "Retrieves an ID namespace association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getIdNamespaceAssociationCmd).Standalone()

	cleanrooms_getIdNamespaceAssociationCmd.Flags().String("id-namespace-association-identifier", "", "The unique identifier of the ID namespace association that you want to retrieve.")
	cleanrooms_getIdNamespaceAssociationCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID namespace association that you want to retrieve.")
	cleanrooms_getIdNamespaceAssociationCmd.MarkFlagRequired("id-namespace-association-identifier")
	cleanrooms_getIdNamespaceAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getIdNamespaceAssociationCmd)
}
