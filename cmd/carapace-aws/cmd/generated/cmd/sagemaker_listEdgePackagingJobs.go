package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listEdgePackagingJobsCmd = &cobra.Command{
	Use:   "list-edge-packaging-jobs",
	Short: "Returns a list of edge packaging jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listEdgePackagingJobsCmd).Standalone()

	sagemaker_listEdgePackagingJobsCmd.Flags().String("creation-time-after", "", "Select jobs where the job was created after specified time.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("creation-time-before", "", "Select jobs where the job was created before specified time.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("last-modified-time-after", "", "Select jobs where the job was updated after specified time.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("last-modified-time-before", "", "Select jobs where the job was updated before specified time.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("max-results", "", "Maximum number of results to select.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("model-name-contains", "", "Filter for jobs where the model name contains this string.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("name-contains", "", "Filter for jobs containing this name in their packaging job name.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to need tokening.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("sort-by", "", "Use to specify what column to sort by.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("sort-order", "", "What direction to sort by.")
	sagemaker_listEdgePackagingJobsCmd.Flags().String("status-equals", "", "The job status to filter for.")
	sagemakerCmd.AddCommand(sagemaker_listEdgePackagingJobsCmd)
}
