package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "Lists the recommendations present in your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listRecommendationsCmd).Standalone()

	sesv2_listRecommendationsCmd.Flags().String("filter", "", "Filters applied when retrieving recommendations.")
	sesv2_listRecommendationsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListRecommendations` to indicate the position in the list of recommendations.")
	sesv2_listRecommendationsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListRecommendations`.")
	sesv2Cmd.AddCommand(sesv2_listRecommendationsCmd)
}
