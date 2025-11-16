package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTrainingJobsForHyperParameterTuningJobCmd = &cobra.Command{
	Use:   "list-training-jobs-for-hyper-parameter-tuning-job",
	Short: "Gets a list of [TrainingJobSummary](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_TrainingJobSummary.html) objects that describe the training jobs that a hyperparameter tuning job launched.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTrainingJobsForHyperParameterTuningJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listTrainingJobsForHyperParameterTuningJobCmd).Standalone()

		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-name", "", "The name of the tuning job whose training jobs you want to list.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("max-results", "", "The maximum number of training jobs to return.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("next-token", "", "If the result of the previous `ListTrainingJobsForHyperParameterTuningJob` request was truncated, the response includes a `NextToken`.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.Flags().String("status-equals", "", "A filter that returns only training jobs with the specified status.")
		sagemaker_listTrainingJobsForHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listTrainingJobsForHyperParameterTuningJobCmd)
}
