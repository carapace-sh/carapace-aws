package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_associateAttributeGroupCmd = &cobra.Command{
	Use:   "associate-attribute-group",
	Short: "Associates an attribute group with an application to augment the application's metadata with the group's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_associateAttributeGroupCmd).Standalone()

	servicecatalogAppregistry_associateAttributeGroupCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_associateAttributeGroupCmd.Flags().String("attribute-group", "", "The name, ID, or ARN of the attribute group that holds the attributes to describe the application.")
	servicecatalogAppregistry_associateAttributeGroupCmd.MarkFlagRequired("application")
	servicecatalogAppregistry_associateAttributeGroupCmd.MarkFlagRequired("attribute-group")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_associateAttributeGroupCmd)
}
