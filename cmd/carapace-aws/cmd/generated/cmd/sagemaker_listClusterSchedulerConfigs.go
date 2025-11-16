package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listClusterSchedulerConfigsCmd = &cobra.Command{
	Use:   "list-cluster-scheduler-configs",
	Short: "List the cluster policy configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listClusterSchedulerConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listClusterSchedulerConfigsCmd).Standalone()

		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("cluster-arn", "", "Filter for ARN of the cluster.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("created-after", "", "Filter for after this creation time.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("created-before", "", "Filter for before this creation time.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("max-results", "", "The maximum number of cluster policies to list.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("name-contains", "", "Filter for name containing this string.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("sort-by", "", "Filter for sorting the list by a given value.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("sort-order", "", "The order of the list.")
		sagemaker_listClusterSchedulerConfigsCmd.Flags().String("status", "", "Filter for status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listClusterSchedulerConfigsCmd)
}
