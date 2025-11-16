package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listProcessingJobsCmd = &cobra.Command{
	Use:   "list-processing-jobs",
	Short: "Lists processing jobs that satisfy various filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listProcessingJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listProcessingJobsCmd).Standalone()

		sagemaker_listProcessingJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only processing jobs created after the specified time.")
		sagemaker_listProcessingJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only processing jobs created after the specified time.")
		sagemaker_listProcessingJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only processing jobs modified after the specified time.")
		sagemaker_listProcessingJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only processing jobs modified before the specified time.")
		sagemaker_listProcessingJobsCmd.Flags().String("max-results", "", "The maximum number of processing jobs to return in the response.")
		sagemaker_listProcessingJobsCmd.Flags().String("name-contains", "", "A string in the processing job name.")
		sagemaker_listProcessingJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListProcessingJobs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listProcessingJobsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listProcessingJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listProcessingJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only processing jobs with a specific status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listProcessingJobsCmd)
}
