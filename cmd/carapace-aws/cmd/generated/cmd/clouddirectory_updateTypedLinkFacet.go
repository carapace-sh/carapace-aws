package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_updateTypedLinkFacetCmd = &cobra.Command{
	Use:   "update-typed-link-facet",
	Short: "Updates a [TypedLinkFacet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_updateTypedLinkFacetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_updateTypedLinkFacetCmd).Standalone()

		clouddirectory_updateTypedLinkFacetCmd.Flags().String("attribute-updates", "", "Attributes update structure.")
		clouddirectory_updateTypedLinkFacetCmd.Flags().String("identity-attribute-order", "", "The order of identity attributes for the facet, from most significant to least significant.")
		clouddirectory_updateTypedLinkFacetCmd.Flags().String("name", "", "The unique name of the typed link facet.")
		clouddirectory_updateTypedLinkFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
		clouddirectory_updateTypedLinkFacetCmd.MarkFlagRequired("attribute-updates")
		clouddirectory_updateTypedLinkFacetCmd.MarkFlagRequired("identity-attribute-order")
		clouddirectory_updateTypedLinkFacetCmd.MarkFlagRequired("name")
		clouddirectory_updateTypedLinkFacetCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_updateTypedLinkFacetCmd)
}
