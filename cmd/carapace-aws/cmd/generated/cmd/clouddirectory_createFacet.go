package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createFacetCmd = &cobra.Command{
	Use:   "create-facet",
	Short: "Creates a new [Facet]() in a schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createFacetCmd).Standalone()

	clouddirectory_createFacetCmd.Flags().String("attributes", "", "The attributes that are associated with the [Facet]().")
	clouddirectory_createFacetCmd.Flags().String("facet-style", "", "There are two different styles that you can define on any given facet, `Static` and `Dynamic`.")
	clouddirectory_createFacetCmd.Flags().String("name", "", "The name of the [Facet](), which is unique for a given schema.")
	clouddirectory_createFacetCmd.Flags().String("object-type", "", "Specifies whether a given object created from this facet is of type node, leaf node, policy or index.")
	clouddirectory_createFacetCmd.Flags().String("schema-arn", "", "The schema ARN in which the new [Facet]() will be created.")
	clouddirectory_createFacetCmd.MarkFlagRequired("name")
	clouddirectory_createFacetCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_createFacetCmd)
}
