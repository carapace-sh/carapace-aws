package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listReviewableHitsCmd = &cobra.Command{
	Use:   "list-reviewable-hits",
	Short: "The `ListReviewableHITs` operation retrieves the HITs with Status equal to Reviewable or Status equal to Reviewing that belong to the Requester calling the operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listReviewableHitsCmd).Standalone()

	mturk_listReviewableHitsCmd.Flags().String("hittype-id", "", "The ID of the HIT type of the HITs to consider for the query.")
	mturk_listReviewableHitsCmd.Flags().String("max-results", "", "Limit the number of results returned.")
	mturk_listReviewableHitsCmd.Flags().String("next-token", "", "Pagination Token")
	mturk_listReviewableHitsCmd.Flags().String("status", "", "Can be either `Reviewable` or `Reviewing`.")
	mturkCmd.AddCommand(mturk_listReviewableHitsCmd)
}
