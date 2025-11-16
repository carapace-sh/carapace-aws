package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getTypedLinkFacetInformationCmd = &cobra.Command{
	Use:   "get-typed-link-facet-information",
	Short: "Returns the identity attribute order for a specific [TypedLinkFacet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getTypedLinkFacetInformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_getTypedLinkFacetInformationCmd).Standalone()

		clouddirectory_getTypedLinkFacetInformationCmd.Flags().String("name", "", "The unique name of the typed link facet.")
		clouddirectory_getTypedLinkFacetInformationCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
		clouddirectory_getTypedLinkFacetInformationCmd.MarkFlagRequired("name")
		clouddirectory_getTypedLinkFacetInformationCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_getTypedLinkFacetInformationCmd)
}
