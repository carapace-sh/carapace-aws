package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeBatchInferenceJobCmd = &cobra.Command{
	Use:   "describe-batch-inference-job",
	Short: "Gets the properties of a batch inference job including name, Amazon Resource Name (ARN), status, input and output configurations, and the ARN of the solution version used to generate the recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeBatchInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeBatchInferenceJobCmd).Standalone()

		personalize_describeBatchInferenceJobCmd.Flags().String("batch-inference-job-arn", "", "The ARN of the batch inference job to describe.")
		personalize_describeBatchInferenceJobCmd.MarkFlagRequired("batch-inference-job-arn")
	})
	personalizeCmd.AddCommand(personalize_describeBatchInferenceJobCmd)
}
