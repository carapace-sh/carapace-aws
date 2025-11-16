package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listBatchInferenceJobsCmd = &cobra.Command{
	Use:   "list-batch-inference-jobs",
	Short: "Gets a list of the batch inference jobs that have been performed off of a solution version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listBatchInferenceJobsCmd).Standalone()

	personalize_listBatchInferenceJobsCmd.Flags().String("max-results", "", "The maximum number of batch inference job results to return in each page.")
	personalize_listBatchInferenceJobsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	personalize_listBatchInferenceJobsCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version from which the batch inference jobs were created.")
	personalizeCmd.AddCommand(personalize_listBatchInferenceJobsCmd)
}
