package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listFacetNamesCmd = &cobra.Command{
	Use:   "list-facet-names",
	Short: "Retrieves the names of facets that exist in a schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listFacetNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listFacetNamesCmd).Standalone()

		clouddirectory_listFacetNamesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listFacetNamesCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listFacetNamesCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) to retrieve facet names from.")
		clouddirectory_listFacetNamesCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listFacetNamesCmd)
}
