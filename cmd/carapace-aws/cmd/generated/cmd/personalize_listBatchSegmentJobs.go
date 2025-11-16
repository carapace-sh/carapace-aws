package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listBatchSegmentJobsCmd = &cobra.Command{
	Use:   "list-batch-segment-jobs",
	Short: "Gets a list of the batch segment jobs that have been performed off of a solution version that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listBatchSegmentJobsCmd).Standalone()

	personalize_listBatchSegmentJobsCmd.Flags().String("max-results", "", "The maximum number of batch segment job results to return in each page.")
	personalize_listBatchSegmentJobsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	personalize_listBatchSegmentJobsCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version that the batch segment jobs used to generate batch segments.")
	personalizeCmd.AddCommand(personalize_listBatchSegmentJobsCmd)
}
