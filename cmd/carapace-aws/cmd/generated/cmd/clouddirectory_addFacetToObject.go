package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_addFacetToObjectCmd = &cobra.Command{
	Use:   "add-facet-to-object",
	Short: "Adds a new [Facet]() to an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_addFacetToObjectCmd).Standalone()

	clouddirectory_addFacetToObjectCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
	clouddirectory_addFacetToObjectCmd.Flags().String("object-attribute-list", "", "Attributes on the facet that you are adding to the object.")
	clouddirectory_addFacetToObjectCmd.Flags().String("object-reference", "", "A reference to the object you are adding the specified facet to.")
	clouddirectory_addFacetToObjectCmd.Flags().String("schema-facet", "", "Identifiers for the facet that you are adding to the object.")
	clouddirectory_addFacetToObjectCmd.MarkFlagRequired("directory-arn")
	clouddirectory_addFacetToObjectCmd.MarkFlagRequired("object-reference")
	clouddirectory_addFacetToObjectCmd.MarkFlagRequired("schema-facet")
	clouddirectoryCmd.AddCommand(clouddirectory_addFacetToObjectCmd)
}
