package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_listCodeReviewsCmd = &cobra.Command{
	Use:   "list-code-reviews",
	Short: "Lists all the code reviews that the customer has created in the past 90 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_listCodeReviewsCmd).Standalone()

	codeguruReviewer_listCodeReviewsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	codeguruReviewer_listCodeReviewsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	codeguruReviewer_listCodeReviewsCmd.Flags().String("provider-types", "", "List of provider types for filtering that needs to be applied before displaying the result.")
	codeguruReviewer_listCodeReviewsCmd.Flags().String("repository-names", "", "List of repository names for filtering that needs to be applied before displaying the result.")
	codeguruReviewer_listCodeReviewsCmd.Flags().String("states", "", "List of states for filtering that needs to be applied before displaying the result.")
	codeguruReviewer_listCodeReviewsCmd.Flags().String("type", "", "The type of code reviews to list in the response.")
	codeguruReviewer_listCodeReviewsCmd.MarkFlagRequired("type")
	codeguruReviewerCmd.AddCommand(codeguruReviewer_listCodeReviewsCmd)
}
