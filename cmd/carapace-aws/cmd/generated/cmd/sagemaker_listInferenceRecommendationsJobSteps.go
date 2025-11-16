package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listInferenceRecommendationsJobStepsCmd = &cobra.Command{
	Use:   "list-inference-recommendations-job-steps",
	Short: "Returns a list of the subtasks for an Inference Recommender job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listInferenceRecommendationsJobStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listInferenceRecommendationsJobStepsCmd).Standalone()

		sagemaker_listInferenceRecommendationsJobStepsCmd.Flags().String("job-name", "", "The name for the Inference Recommender job.")
		sagemaker_listInferenceRecommendationsJobStepsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		sagemaker_listInferenceRecommendationsJobStepsCmd.Flags().String("next-token", "", "A token that you can specify to return more results from the list.")
		sagemaker_listInferenceRecommendationsJobStepsCmd.Flags().String("status", "", "A filter to return benchmarks of a specified status.")
		sagemaker_listInferenceRecommendationsJobStepsCmd.Flags().String("step-type", "", "A filter to return details about the specified type of subtask.")
		sagemaker_listInferenceRecommendationsJobStepsCmd.MarkFlagRequired("job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listInferenceRecommendationsJobStepsCmd)
}
