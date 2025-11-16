package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listFeaturedResultsSetsCmd = &cobra.Command{
	Use:   "list-featured-results-sets",
	Short: "Lists all your sets of featured results for a given index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listFeaturedResultsSetsCmd).Standalone()

	kendra_listFeaturedResultsSetsCmd.Flags().String("index-id", "", "The identifier of the index used for featuring results.")
	kendra_listFeaturedResultsSetsCmd.Flags().String("max-results", "", "The maximum number of featured results sets to return.")
	kendra_listFeaturedResultsSetsCmd.Flags().String("next-token", "", "If the response is truncated, Amazon Kendra returns a pagination token in the response.")
	kendra_listFeaturedResultsSetsCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listFeaturedResultsSetsCmd)
}
