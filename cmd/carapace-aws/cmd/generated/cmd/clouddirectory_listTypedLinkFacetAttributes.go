package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listTypedLinkFacetAttributesCmd = &cobra.Command{
	Use:   "list-typed-link-facet-attributes",
	Short: "Returns a paginated list of all attribute definitions for a particular [TypedLinkFacet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listTypedLinkFacetAttributesCmd).Standalone()

	clouddirectory_listTypedLinkFacetAttributesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listTypedLinkFacetAttributesCmd.Flags().String("name", "", "The unique name of the typed link facet.")
	clouddirectory_listTypedLinkFacetAttributesCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listTypedLinkFacetAttributesCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
	clouddirectory_listTypedLinkFacetAttributesCmd.MarkFlagRequired("name")
	clouddirectory_listTypedLinkFacetAttributesCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_listTypedLinkFacetAttributesCmd)
}
