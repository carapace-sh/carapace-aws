package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_searchRelatedItemsCmd = &cobra.Command{
	Use:   "search-related-items",
	Short: "Searches for related items that are associated with a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_searchRelatedItemsCmd).Standalone()

	connectcases_searchRelatedItemsCmd.Flags().String("case-id", "", "A unique identifier of the case.")
	connectcases_searchRelatedItemsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_searchRelatedItemsCmd.Flags().String("filters", "", "The list of types of related items and their parameters to use for filtering.")
	connectcases_searchRelatedItemsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_searchRelatedItemsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_searchRelatedItemsCmd.MarkFlagRequired("case-id")
	connectcases_searchRelatedItemsCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_searchRelatedItemsCmd)
}
