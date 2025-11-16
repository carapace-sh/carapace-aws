package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createBatchInferenceJobCmd = &cobra.Command{
	Use:   "create-batch-inference-job",
	Short: "Generates batch recommendations based on a list of items or users stored in Amazon S3 and exports the recommendations to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createBatchInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_createBatchInferenceJobCmd).Standalone()

		personalize_createBatchInferenceJobCmd.Flags().String("batch-inference-job-config", "", "The configuration details of a batch inference job.")
		personalize_createBatchInferenceJobCmd.Flags().String("batch-inference-job-mode", "", "The mode of the batch inference job.")
		personalize_createBatchInferenceJobCmd.Flags().String("filter-arn", "", "The ARN of the filter to apply to the batch inference job.")
		personalize_createBatchInferenceJobCmd.Flags().String("job-input", "", "The Amazon S3 path that leads to the input file to base your recommendations on.")
		personalize_createBatchInferenceJobCmd.Flags().String("job-name", "", "The name of the batch inference job to create.")
		personalize_createBatchInferenceJobCmd.Flags().String("job-output", "", "The path to the Amazon S3 bucket where the job's output will be stored.")
		personalize_createBatchInferenceJobCmd.Flags().String("num-results", "", "The number of recommendations to retrieve.")
		personalize_createBatchInferenceJobCmd.Flags().String("role-arn", "", "The ARN of the Amazon Identity and Access Management role that has permissions to read and write to your input and output Amazon S3 buckets respectively.")
		personalize_createBatchInferenceJobCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version that will be used to generate the batch inference recommendations.")
		personalize_createBatchInferenceJobCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the batch inference job.")
		personalize_createBatchInferenceJobCmd.Flags().String("theme-generation-config", "", "For theme generation jobs, specify the name of the column in your Items dataset that contains each item's name.")
		personalize_createBatchInferenceJobCmd.MarkFlagRequired("job-input")
		personalize_createBatchInferenceJobCmd.MarkFlagRequired("job-name")
		personalize_createBatchInferenceJobCmd.MarkFlagRequired("job-output")
		personalize_createBatchInferenceJobCmd.MarkFlagRequired("role-arn")
		personalize_createBatchInferenceJobCmd.MarkFlagRequired("solution-version-arn")
	})
	personalizeCmd.AddCommand(personalize_createBatchInferenceJobCmd)
}
