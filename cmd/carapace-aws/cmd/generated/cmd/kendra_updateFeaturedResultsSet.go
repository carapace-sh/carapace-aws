package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateFeaturedResultsSetCmd = &cobra.Command{
	Use:   "update-featured-results-set",
	Short: "Updates a set of featured results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateFeaturedResultsSetCmd).Standalone()

	kendra_updateFeaturedResultsSetCmd.Flags().String("description", "", "A new description for the set of featured results.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("featured-documents", "", "A list of document IDs for the documents you want to feature at the top of the search results page.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("featured-results-set-id", "", "The identifier of the set of featured results that you want to update.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("featured-results-set-name", "", "A new name for the set of featured results.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("index-id", "", "The identifier of the index used for featuring results.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("query-texts", "", "A list of queries for featuring results.")
	kendra_updateFeaturedResultsSetCmd.Flags().String("status", "", "You can set the status to `ACTIVE` or `INACTIVE`.")
	kendra_updateFeaturedResultsSetCmd.MarkFlagRequired("featured-results-set-id")
	kendra_updateFeaturedResultsSetCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_updateFeaturedResultsSetCmd)
}
