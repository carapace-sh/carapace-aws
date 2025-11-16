package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_deleteTypedLinkFacetCmd = &cobra.Command{
	Use:   "delete-typed-link-facet",
	Short: "Deletes a [TypedLinkFacet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_deleteTypedLinkFacetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_deleteTypedLinkFacetCmd).Standalone()

		clouddirectory_deleteTypedLinkFacetCmd.Flags().String("name", "", "The unique name of the typed link facet.")
		clouddirectory_deleteTypedLinkFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
		clouddirectory_deleteTypedLinkFacetCmd.MarkFlagRequired("name")
		clouddirectory_deleteTypedLinkFacetCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_deleteTypedLinkFacetCmd)
}
