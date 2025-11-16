package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_searchQuickResponsesCmd = &cobra.Command{
	Use:   "search-quick-responses",
	Short: "Searches existing Wisdom quick responses in a Wisdom knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_searchQuickResponsesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_searchQuickResponsesCmd).Standalone()

		wisdom_searchQuickResponsesCmd.Flags().String("attributes", "", "The [user-defined Amazon Connect contact attributes](https://docs.aws.amazon.com/connect/latest/adminguide/connect-attrib-list.html#user-defined-attributes) to be resolved when search results are returned.")
		wisdom_searchQuickResponsesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_searchQuickResponsesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		wisdom_searchQuickResponsesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		wisdom_searchQuickResponsesCmd.Flags().String("search-expression", "", "The search expression for querying the quick response.")
		wisdom_searchQuickResponsesCmd.MarkFlagRequired("knowledge-base-id")
		wisdom_searchQuickResponsesCmd.MarkFlagRequired("search-expression")
	})
	wisdomCmd.AddCommand(wisdom_searchQuickResponsesCmd)
}
