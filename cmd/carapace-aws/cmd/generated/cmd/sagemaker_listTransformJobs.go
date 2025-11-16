package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTransformJobsCmd = &cobra.Command{
	Use:   "list-transform-jobs",
	Short: "Lists transform jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTransformJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listTransformJobsCmd).Standalone()

		sagemaker_listTransformJobsCmd.Flags().String("creation-time-after", "", "A filter that returns only transform jobs created after the specified time.")
		sagemaker_listTransformJobsCmd.Flags().String("creation-time-before", "", "A filter that returns only transform jobs created before the specified time.")
		sagemaker_listTransformJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only transform jobs modified after the specified time.")
		sagemaker_listTransformJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only transform jobs modified before the specified time.")
		sagemaker_listTransformJobsCmd.Flags().String("max-results", "", "The maximum number of transform jobs to return in the response.")
		sagemaker_listTransformJobsCmd.Flags().String("name-contains", "", "A string in the transform job name.")
		sagemaker_listTransformJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListTransformJobs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listTransformJobsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listTransformJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listTransformJobsCmd.Flags().String("status-equals", "", "A filter that retrieves only transform jobs with a specific status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listTransformJobsCmd)
}
