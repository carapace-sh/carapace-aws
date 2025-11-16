package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listOptimizationJobsCmd = &cobra.Command{
	Use:   "list-optimization-jobs",
	Short: "Lists the optimization jobs in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listOptimizationJobsCmd).Standalone()

	sagemaker_listOptimizationJobsCmd.Flags().String("creation-time-after", "", "Filters the results to only those optimization jobs that were created after the specified time.")
	sagemaker_listOptimizationJobsCmd.Flags().String("creation-time-before", "", "Filters the results to only those optimization jobs that were created before the specified time.")
	sagemaker_listOptimizationJobsCmd.Flags().String("last-modified-time-after", "", "Filters the results to only those optimization jobs that were updated after the specified time.")
	sagemaker_listOptimizationJobsCmd.Flags().String("last-modified-time-before", "", "Filters the results to only those optimization jobs that were updated before the specified time.")
	sagemaker_listOptimizationJobsCmd.Flags().String("max-results", "", "The maximum number of optimization jobs to return in the response.")
	sagemaker_listOptimizationJobsCmd.Flags().String("name-contains", "", "Filters the results to only those optimization jobs with a name that contains the specified string.")
	sagemaker_listOptimizationJobsCmd.Flags().String("next-token", "", "A token that you use to get the next set of results following a truncated response.")
	sagemaker_listOptimizationJobsCmd.Flags().String("optimization-contains", "", "Filters the results to only those optimization jobs that apply the specified optimization techniques.")
	sagemaker_listOptimizationJobsCmd.Flags().String("sort-by", "", "The field by which to sort the optimization jobs in the response.")
	sagemaker_listOptimizationJobsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemaker_listOptimizationJobsCmd.Flags().String("status-equals", "", "Filters the results to only those optimization jobs with the specified status.")
	sagemakerCmd.AddCommand(sagemaker_listOptimizationJobsCmd)
}
