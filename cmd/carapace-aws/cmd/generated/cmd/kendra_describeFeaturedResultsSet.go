package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeFeaturedResultsSetCmd = &cobra.Command{
	Use:   "describe-featured-results-set",
	Short: "Gets information about a set of featured results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeFeaturedResultsSetCmd).Standalone()

	kendra_describeFeaturedResultsSetCmd.Flags().String("featured-results-set-id", "", "The identifier of the set of featured results that you want to get information on.")
	kendra_describeFeaturedResultsSetCmd.Flags().String("index-id", "", "The identifier of the index used for featuring results.")
	kendra_describeFeaturedResultsSetCmd.MarkFlagRequired("featured-results-set-id")
	kendra_describeFeaturedResultsSetCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_describeFeaturedResultsSetCmd)
}
