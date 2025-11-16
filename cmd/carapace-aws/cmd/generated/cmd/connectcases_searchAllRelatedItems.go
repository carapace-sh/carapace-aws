package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_searchAllRelatedItemsCmd = &cobra.Command{
	Use:   "search-all-related-items",
	Short: "Searches for related items across all cases within a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_searchAllRelatedItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_searchAllRelatedItemsCmd).Standalone()

		connectcases_searchAllRelatedItemsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_searchAllRelatedItemsCmd.Flags().String("filters", "", "The list of types of related items and their parameters to use for filtering.")
		connectcases_searchAllRelatedItemsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connectcases_searchAllRelatedItemsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connectcases_searchAllRelatedItemsCmd.Flags().String("sorts", "", "A structured set of sort terms to specify the order in which related items should be returned.")
		connectcases_searchAllRelatedItemsCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_searchAllRelatedItemsCmd)
}
