package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listHyperParameterTuningJobsCmd = &cobra.Command{
	Use:   "list-hyper-parameter-tuning-jobs",
	Short: "Gets a list of [HyperParameterTuningJobSummary](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobSummary.html) objects that describe the hyperparameter tuning jobs launched in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listHyperParameterTuningJobsCmd).Standalone()

	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only tuning jobs that were created after the specified time.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only tuning jobs that were created before the specified time.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only tuning jobs that were modified after the specified time.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only tuning jobs that were modified before the specified time.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("max-results", "", "The maximum number of tuning jobs to return.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("name-contains", "", "A string in the tuning job name.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListHyperParameterTuningJobs` request was truncated, the response includes a `NextToken`.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("sort-by", "", "The field to sort results by.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemaker_listHyperParameterTuningJobsCmd.Flags().String("status-equals", "", "A filter that returns only tuning jobs with the specified status.")
	sagemakerCmd.AddCommand(sagemaker_listHyperParameterTuningJobsCmd)
}
