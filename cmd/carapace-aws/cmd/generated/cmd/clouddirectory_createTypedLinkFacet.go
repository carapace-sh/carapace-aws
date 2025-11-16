package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createTypedLinkFacetCmd = &cobra.Command{
	Use:   "create-typed-link-facet",
	Short: "Creates a [TypedLinkFacet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createTypedLinkFacetCmd).Standalone()

	clouddirectory_createTypedLinkFacetCmd.Flags().String("facet", "", "[Facet]() structure that is associated with the typed link facet.")
	clouddirectory_createTypedLinkFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
	clouddirectory_createTypedLinkFacetCmd.MarkFlagRequired("facet")
	clouddirectory_createTypedLinkFacetCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_createTypedLinkFacetCmd)
}
