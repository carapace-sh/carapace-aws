package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createIdNamespaceAssociationCmd = &cobra.Command{
	Use:   "create-id-namespace-association",
	Short: "Creates an ID namespace association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createIdNamespaceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createIdNamespaceAssociationCmd).Standalone()

		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("description", "", "The description of the ID namespace association.")
		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("id-mapping-config", "", "The configuration settings for the ID mapping table.")
		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("input-reference-config", "", "The input reference configuration needed to create the ID namespace association.")
		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID namespace association.")
		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("name", "", "The name for the ID namespace association.")
		cleanrooms_createIdNamespaceAssociationCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
		cleanrooms_createIdNamespaceAssociationCmd.MarkFlagRequired("input-reference-config")
		cleanrooms_createIdNamespaceAssociationCmd.MarkFlagRequired("membership-identifier")
		cleanrooms_createIdNamespaceAssociationCmd.MarkFlagRequired("name")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createIdNamespaceAssociationCmd)
}
