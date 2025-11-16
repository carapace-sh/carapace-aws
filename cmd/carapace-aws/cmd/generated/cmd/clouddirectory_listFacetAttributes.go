package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listFacetAttributesCmd = &cobra.Command{
	Use:   "list-facet-attributes",
	Short: "Retrieves attributes attached to the facet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listFacetAttributesCmd).Standalone()

	clouddirectory_listFacetAttributesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listFacetAttributesCmd.Flags().String("name", "", "The name of the facet whose attributes will be retrieved.")
	clouddirectory_listFacetAttributesCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listFacetAttributesCmd.Flags().String("schema-arn", "", "The ARN of the schema where the facet resides.")
	clouddirectory_listFacetAttributesCmd.MarkFlagRequired("name")
	clouddirectory_listFacetAttributesCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_listFacetAttributesCmd)
}
