package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listInferenceRecommendationsJobsCmd = &cobra.Command{
	Use:   "list-inference-recommendations-jobs",
	Short: "Lists recommendation jobs that satisfy various filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listInferenceRecommendationsJobsCmd).Standalone()

	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only jobs created after the specified time (timestamp).")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only jobs created before the specified time (timestamp).")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only jobs that were last modified after the specified time (timestamp).")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only jobs that were last modified before the specified time (timestamp).")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("max-results", "", "The maximum number of recommendations to return in the response.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("model-name-equals", "", "A filter that returns only jobs that were created for this model.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("model-package-version-arn-equals", "", "A filter that returns only jobs that were created for this versioned model package.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("name-contains", "", "A string in the job name.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("next-token", "", "If the response to a previous `ListInferenceRecommendationsJobsRequest` request was truncated, the response includes a `NextToken`.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	sagemaker_listInferenceRecommendationsJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only inference recommendations jobs with a specific status.")
	sagemakerCmd.AddCommand(sagemaker_listInferenceRecommendationsJobsCmd)
}
