package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_updateAttributeGroupCmd = &cobra.Command{
	Use:   "update-attribute-group",
	Short: "Updates an existing attribute group with new details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_updateAttributeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_updateAttributeGroupCmd).Standalone()

		servicecatalogAppregistry_updateAttributeGroupCmd.Flags().String("attribute-group", "", "The name, ID, or ARN of the attribute group that holds the attributes to describe the application.")
		servicecatalogAppregistry_updateAttributeGroupCmd.Flags().String("attributes", "", "A JSON string in the form of nested key-value pairs that represent the attributes in the group and describes an application and its components.")
		servicecatalogAppregistry_updateAttributeGroupCmd.Flags().String("description", "", "The description of the attribute group that the user provides.")
		servicecatalogAppregistry_updateAttributeGroupCmd.Flags().String("name", "", "Deprecated: The new name of the attribute group.")
		servicecatalogAppregistry_updateAttributeGroupCmd.MarkFlagRequired("attribute-group")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_updateAttributeGroupCmd)
}
