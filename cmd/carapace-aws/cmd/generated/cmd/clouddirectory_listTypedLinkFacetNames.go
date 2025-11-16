package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listTypedLinkFacetNamesCmd = &cobra.Command{
	Use:   "list-typed-link-facet-names",
	Short: "Returns a paginated list of `TypedLink` facet names for a particular schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listTypedLinkFacetNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listTypedLinkFacetNamesCmd).Standalone()

		clouddirectory_listTypedLinkFacetNamesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listTypedLinkFacetNamesCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listTypedLinkFacetNamesCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the schema.")
		clouddirectory_listTypedLinkFacetNamesCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listTypedLinkFacetNamesCmd)
}
