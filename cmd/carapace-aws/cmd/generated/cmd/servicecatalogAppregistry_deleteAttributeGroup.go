package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_deleteAttributeGroupCmd = &cobra.Command{
	Use:   "delete-attribute-group",
	Short: "Deletes an attribute group, specified either by its attribute group ID, name, or ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_deleteAttributeGroupCmd).Standalone()

	servicecatalogAppregistry_deleteAttributeGroupCmd.Flags().String("attribute-group", "", "The name, ID, or ARN of the attribute group that holds the attributes to describe the application.")
	servicecatalogAppregistry_deleteAttributeGroupCmd.MarkFlagRequired("attribute-group")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_deleteAttributeGroupCmd)
}
