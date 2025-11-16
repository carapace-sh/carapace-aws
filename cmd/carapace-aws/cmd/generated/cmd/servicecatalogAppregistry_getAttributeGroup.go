package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_getAttributeGroupCmd = &cobra.Command{
	Use:   "get-attribute-group",
	Short: "Retrieves an attribute group by its ARN, ID, or name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_getAttributeGroupCmd).Standalone()

	servicecatalogAppregistry_getAttributeGroupCmd.Flags().String("attribute-group", "", "The name, ID, or ARN of the attribute group that holds the attributes to describe the application.")
	servicecatalogAppregistry_getAttributeGroupCmd.MarkFlagRequired("attribute-group")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_getAttributeGroupCmd)
}
