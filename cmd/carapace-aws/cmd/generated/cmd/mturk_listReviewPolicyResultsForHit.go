package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listReviewPolicyResultsForHitCmd = &cobra.Command{
	Use:   "list-review-policy-results-for-hit",
	Short: "The `ListReviewPolicyResultsForHIT` operation retrieves the computed results and the actions taken in the course of executing your Review Policies for a given HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listReviewPolicyResultsForHitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listReviewPolicyResultsForHitCmd).Standalone()

		mturk_listReviewPolicyResultsForHitCmd.Flags().String("hitid", "", "The unique identifier of the HIT to retrieve review results for.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().String("max-results", "", "Limit the number of results returned.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().String("next-token", "", "Pagination token")
		mturk_listReviewPolicyResultsForHitCmd.Flags().Bool("no-retrieve-actions", false, "Specify if the operation should retrieve a list of the actions taken executing the Review Policies and their outcomes.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().Bool("no-retrieve-results", false, "Specify if the operation should retrieve a list of the results computed by the Review Policies.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().String("policy-levels", "", "The Policy Level(s) to retrieve review results for - HIT or Assignment.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().Bool("retrieve-actions", false, "Specify if the operation should retrieve a list of the actions taken executing the Review Policies and their outcomes.")
		mturk_listReviewPolicyResultsForHitCmd.Flags().Bool("retrieve-results", false, "Specify if the operation should retrieve a list of the results computed by the Review Policies.")
		mturk_listReviewPolicyResultsForHitCmd.MarkFlagRequired("hitid")
		mturk_listReviewPolicyResultsForHitCmd.Flag("no-retrieve-actions").Hidden = true
		mturk_listReviewPolicyResultsForHitCmd.Flag("no-retrieve-results").Hidden = true
	})
	mturkCmd.AddCommand(mturk_listReviewPolicyResultsForHitCmd)
}
