package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createFeaturedResultsSetCmd = &cobra.Command{
	Use:   "create-featured-results-set",
	Short: "Creates a set of featured results to display at the top of the search results page.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createFeaturedResultsSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createFeaturedResultsSetCmd).Standalone()

		kendra_createFeaturedResultsSetCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a set of featured results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("description", "", "A description for the set of featured results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("featured-documents", "", "A list of document IDs for the documents you want to feature at the top of the search results page.")
		kendra_createFeaturedResultsSetCmd.Flags().String("featured-results-set-name", "", "A name for the set of featured results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("index-id", "", "The identifier of the index that you want to use for featuring results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("query-texts", "", "A list of queries for featuring results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("status", "", "The current status of the set of featured results.")
		kendra_createFeaturedResultsSetCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the featured results set.")
		kendra_createFeaturedResultsSetCmd.MarkFlagRequired("featured-results-set-name")
		kendra_createFeaturedResultsSetCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_createFeaturedResultsSetCmd)
}
