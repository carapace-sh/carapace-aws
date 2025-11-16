package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_createAttributeGroupCmd = &cobra.Command{
	Use:   "create-attribute-group",
	Short: "Creates a new attribute group as a container for user-defined attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_createAttributeGroupCmd).Standalone()

	servicecatalogAppregistry_createAttributeGroupCmd.Flags().String("attributes", "", "A JSON string in the form of nested key-value pairs that represent the attributes in the group and describes an application and its components.")
	servicecatalogAppregistry_createAttributeGroupCmd.Flags().String("client-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalogAppregistry_createAttributeGroupCmd.Flags().String("description", "", "The description of the attribute group that the user provides.")
	servicecatalogAppregistry_createAttributeGroupCmd.Flags().String("name", "", "The name of the attribute group.")
	servicecatalogAppregistry_createAttributeGroupCmd.Flags().String("tags", "", "Key-value pairs you can use to associate with the attribute group.")
	servicecatalogAppregistry_createAttributeGroupCmd.MarkFlagRequired("attributes")
	servicecatalogAppregistry_createAttributeGroupCmd.MarkFlagRequired("client-token")
	servicecatalogAppregistry_createAttributeGroupCmd.MarkFlagRequired("name")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_createAttributeGroupCmd)
}
