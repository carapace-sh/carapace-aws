package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTrainingJobsCmd = &cobra.Command{
	Use:   "list-training-jobs",
	Short: "Lists training jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTrainingJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listTrainingJobsCmd).Standalone()

		sagemaker_listTrainingJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only training jobs created after the specified time (timestamp).")
		sagemaker_listTrainingJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only training jobs created before the specified time (timestamp).")
		sagemaker_listTrainingJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only training jobs modified after the specified time (timestamp).")
		sagemaker_listTrainingJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only training jobs modified before the specified time (timestamp).")
		sagemaker_listTrainingJobsCmd.Flags().String("max-results", "", "The maximum number of training jobs to return in the response.")
		sagemaker_listTrainingJobsCmd.Flags().String("name-contains", "", "A string in the training job name.")
		sagemaker_listTrainingJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListTrainingJobs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listTrainingJobsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listTrainingJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listTrainingJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only training jobs with a specific status.")
		sagemaker_listTrainingJobsCmd.Flags().String("training-plan-arn-equals", "", "The Amazon Resource Name (ARN); of the training plan to filter training jobs by.")
		sagemaker_listTrainingJobsCmd.Flags().String("warm-pool-status-equals", "", "A filter that retrieves only training jobs with a specific warm pool status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listTrainingJobsCmd)
}
