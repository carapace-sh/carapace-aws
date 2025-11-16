package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelCardExportJobsCmd = &cobra.Command{
	Use:   "list-model-card-export-jobs",
	Short: "List the export jobs for the Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelCardExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listModelCardExportJobsCmd).Standalone()

		sagemaker_listModelCardExportJobsCmd.Flags().String("creation-time-after", "", "Only list model card export jobs that were created after the time specified.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("creation-time-before", "", "Only list model card export jobs that were created before the time specified.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("max-results", "", "The maximum number of model card export jobs to list.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("model-card-export-job-name-contains", "", "Only list model card export jobs with names that contain the specified string.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("model-card-name", "", "List export jobs for the model card with the specified name.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("model-card-version", "", "List export jobs for the model card with the specified version.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("next-token", "", "If the response to a previous `ListModelCardExportJobs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("sort-by", "", "Sort model card export jobs by either name or creation time.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("sort-order", "", "Sort model card export jobs by ascending or descending order.")
		sagemaker_listModelCardExportJobsCmd.Flags().String("status-equals", "", "Only list model card export jobs with the specified status.")
		sagemaker_listModelCardExportJobsCmd.MarkFlagRequired("model-card-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listModelCardExportJobsCmd)
}
