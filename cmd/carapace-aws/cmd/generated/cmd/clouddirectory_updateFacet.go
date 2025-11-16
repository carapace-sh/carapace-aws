package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_updateFacetCmd = &cobra.Command{
	Use:   "update-facet",
	Short: "Does the following:\n\n1. Adds new `Attributes`, `Rules`, or `ObjectTypes`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_updateFacetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_updateFacetCmd).Standalone()

		clouddirectory_updateFacetCmd.Flags().String("attribute-updates", "", "List of attributes that need to be updated in a given schema [Facet]().")
		clouddirectory_updateFacetCmd.Flags().String("name", "", "The name of the facet.")
		clouddirectory_updateFacetCmd.Flags().String("object-type", "", "The object type that is associated with the facet.")
		clouddirectory_updateFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Facet]().")
		clouddirectory_updateFacetCmd.MarkFlagRequired("name")
		clouddirectory_updateFacetCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_updateFacetCmd)
}
