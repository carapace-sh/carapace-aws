package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateIdNamespaceAssociationCmd = &cobra.Command{
	Use:   "update-id-namespace-association",
	Short: "Provides the details that are necessary to update an ID namespace association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateIdNamespaceAssociationCmd).Standalone()

	cleanrooms_updateIdNamespaceAssociationCmd.Flags().String("description", "", "A new description for the ID namespace association.")
	cleanrooms_updateIdNamespaceAssociationCmd.Flags().String("id-mapping-config", "", "The configuration settings for the ID mapping table.")
	cleanrooms_updateIdNamespaceAssociationCmd.Flags().String("id-namespace-association-identifier", "", "The unique identifier of the ID namespace association that you want to update.")
	cleanrooms_updateIdNamespaceAssociationCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID namespace association that you want to update.")
	cleanrooms_updateIdNamespaceAssociationCmd.Flags().String("name", "", "A new name for the ID namespace association.")
	cleanrooms_updateIdNamespaceAssociationCmd.MarkFlagRequired("id-namespace-association-identifier")
	cleanrooms_updateIdNamespaceAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_updateIdNamespaceAssociationCmd)
}
