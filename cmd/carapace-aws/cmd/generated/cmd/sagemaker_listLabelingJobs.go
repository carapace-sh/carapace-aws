package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listLabelingJobsCmd = &cobra.Command{
	Use:   "list-labeling-jobs",
	Short: "Gets a list of labeling jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listLabelingJobsCmd).Standalone()

	sagemaker_listLabelingJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only labeling jobs created after the specified time (timestamp).")
	sagemaker_listLabelingJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only labeling jobs created before the specified time (timestamp).")
	sagemaker_listLabelingJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only labeling jobs modified after the specified time (timestamp).")
	sagemaker_listLabelingJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only labeling jobs modified before the specified time (timestamp).")
	sagemaker_listLabelingJobsCmd.Flags().String("max-results", "", "The maximum number of labeling jobs to return in each page of the response.")
	sagemaker_listLabelingJobsCmd.Flags().String("name-contains", "", "A string in the labeling job name.")
	sagemaker_listLabelingJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListLabelingJobs` request was truncated, the response includes a `NextToken`.")
	sagemaker_listLabelingJobsCmd.Flags().String("sort-by", "", "The field to sort results by.")
	sagemaker_listLabelingJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemaker_listLabelingJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only labeling jobs with a specific status.")
	sagemakerCmd.AddCommand(sagemaker_listLabelingJobsCmd)
}
