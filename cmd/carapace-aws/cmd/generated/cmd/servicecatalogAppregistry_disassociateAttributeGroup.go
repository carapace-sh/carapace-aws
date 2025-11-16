package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_disassociateAttributeGroupCmd = &cobra.Command{
	Use:   "disassociate-attribute-group",
	Short: "Disassociates an attribute group from an application to remove the extra attributes contained in the attribute group from the application's metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_disassociateAttributeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_disassociateAttributeGroupCmd).Standalone()

		servicecatalogAppregistry_disassociateAttributeGroupCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
		servicecatalogAppregistry_disassociateAttributeGroupCmd.Flags().String("attribute-group", "", "The name, ID, or ARN of the attribute group that holds the attributes to describe the application.")
		servicecatalogAppregistry_disassociateAttributeGroupCmd.MarkFlagRequired("application")
		servicecatalogAppregistry_disassociateAttributeGroupCmd.MarkFlagRequired("attribute-group")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_disassociateAttributeGroupCmd)
}
