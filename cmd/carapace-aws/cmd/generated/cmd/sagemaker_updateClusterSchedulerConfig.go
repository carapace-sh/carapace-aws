package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateClusterSchedulerConfigCmd = &cobra.Command{
	Use:   "update-cluster-scheduler-config",
	Short: "Update the cluster policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateClusterSchedulerConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateClusterSchedulerConfigCmd).Standalone()

		sagemaker_updateClusterSchedulerConfigCmd.Flags().String("cluster-scheduler-config-id", "", "ID of the cluster policy.")
		sagemaker_updateClusterSchedulerConfigCmd.Flags().String("description", "", "Description of the cluster policy.")
		sagemaker_updateClusterSchedulerConfigCmd.Flags().String("scheduler-config", "", "Cluster policy configuration.")
		sagemaker_updateClusterSchedulerConfigCmd.Flags().String("target-version", "", "Target version.")
		sagemaker_updateClusterSchedulerConfigCmd.MarkFlagRequired("cluster-scheduler-config-id")
		sagemaker_updateClusterSchedulerConfigCmd.MarkFlagRequired("target-version")
	})
	sagemakerCmd.AddCommand(sagemaker_updateClusterSchedulerConfigCmd)
}
