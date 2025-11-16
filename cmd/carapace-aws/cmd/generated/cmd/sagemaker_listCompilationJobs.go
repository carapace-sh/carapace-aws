package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listCompilationJobsCmd = &cobra.Command{
	Use:   "list-compilation-jobs",
	Short: "Lists model compilation jobs that satisfy various filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listCompilationJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listCompilationJobsCmd).Standalone()

		sagemaker_listCompilationJobsCmd.Flags().String("creation-time-after", "", "A filter that returns the model compilation jobs that were created after a specified time.")
		sagemaker_listCompilationJobsCmd.Flags().String("creation-time-before", "", "A filter that returns the model compilation jobs that were created before a specified time.")
		sagemaker_listCompilationJobsCmd.Flags().String("last-modified-time-after", "", "A filter that returns the model compilation jobs that were modified after a specified time.")
		sagemaker_listCompilationJobsCmd.Flags().String("last-modified-time-before", "", "A filter that returns the model compilation jobs that were modified before a specified time.")
		sagemaker_listCompilationJobsCmd.Flags().String("max-results", "", "The maximum number of model compilation jobs to return in the response.")
		sagemaker_listCompilationJobsCmd.Flags().String("name-contains", "", "A filter that returns the model compilation jobs whose name contains a specified string.")
		sagemaker_listCompilationJobsCmd.Flags().String("next-token", "", "If the result of the previous `ListCompilationJobs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listCompilationJobsCmd.Flags().String("sort-by", "", "The field by which to sort results.")
		sagemaker_listCompilationJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listCompilationJobsCmd.Flags().String("status-equals", "", "A filter that retrieves model compilation jobs with a specific `CompilationJobStatus` status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listCompilationJobsCmd)
}
