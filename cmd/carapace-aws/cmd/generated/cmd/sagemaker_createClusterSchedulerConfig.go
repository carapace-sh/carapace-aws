package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createClusterSchedulerConfigCmd = &cobra.Command{
	Use:   "create-cluster-scheduler-config",
	Short: "Create cluster policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createClusterSchedulerConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createClusterSchedulerConfigCmd).Standalone()

		sagemaker_createClusterSchedulerConfigCmd.Flags().String("cluster-arn", "", "ARN of the cluster.")
		sagemaker_createClusterSchedulerConfigCmd.Flags().String("description", "", "Description of the cluster policy.")
		sagemaker_createClusterSchedulerConfigCmd.Flags().String("name", "", "Name for the cluster policy.")
		sagemaker_createClusterSchedulerConfigCmd.Flags().String("scheduler-config", "", "Configuration about the monitoring schedule.")
		sagemaker_createClusterSchedulerConfigCmd.Flags().String("tags", "", "Tags of the cluster policy.")
		sagemaker_createClusterSchedulerConfigCmd.MarkFlagRequired("cluster-arn")
		sagemaker_createClusterSchedulerConfigCmd.MarkFlagRequired("name")
		sagemaker_createClusterSchedulerConfigCmd.MarkFlagRequired("scheduler-config")
	})
	sagemakerCmd.AddCommand(sagemaker_createClusterSchedulerConfigCmd)
}
