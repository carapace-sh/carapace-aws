package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_batchDeleteFeaturedResultsSetCmd = &cobra.Command{
	Use:   "batch-delete-featured-results-set",
	Short: "Removes one or more sets of featured results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_batchDeleteFeaturedResultsSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_batchDeleteFeaturedResultsSetCmd).Standalone()

		kendra_batchDeleteFeaturedResultsSetCmd.Flags().String("featured-results-set-ids", "", "The identifiers of the featured results sets that you want to delete.")
		kendra_batchDeleteFeaturedResultsSetCmd.Flags().String("index-id", "", "The identifier of the index used for featuring results.")
		kendra_batchDeleteFeaturedResultsSetCmd.MarkFlagRequired("featured-results-set-ids")
		kendra_batchDeleteFeaturedResultsSetCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_batchDeleteFeaturedResultsSetCmd)
}
