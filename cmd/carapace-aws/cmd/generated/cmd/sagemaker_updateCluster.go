package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Updates a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateClusterCmd).Standalone()

		sagemaker_updateClusterCmd.Flags().String("auto-scaling", "", "Updates the autoscaling configuration for the cluster.")
		sagemaker_updateClusterCmd.Flags().String("cluster-name", "", "Specify the name of the SageMaker HyperPod cluster you want to update.")
		sagemaker_updateClusterCmd.Flags().String("cluster-role", "", "The Amazon Resource Name (ARN) of the IAM role that HyperPod assumes for cluster autoscaling operations.")
		sagemaker_updateClusterCmd.Flags().String("instance-groups", "", "Specify the instance groups to update.")
		sagemaker_updateClusterCmd.Flags().String("instance-groups-to-delete", "", "Specify the names of the instance groups to delete.")
		sagemaker_updateClusterCmd.Flags().String("node-provisioning-mode", "", "Determines how instance provisioning is handled during cluster operations.")
		sagemaker_updateClusterCmd.Flags().String("node-recovery", "", "The node recovery mode to be applied to the SageMaker HyperPod cluster.")
		sagemaker_updateClusterCmd.Flags().String("restricted-instance-groups", "", "The specialized instance groups for training models like Amazon Nova to be created in the SageMaker HyperPod cluster.")
		sagemaker_updateClusterCmd.Flags().String("tiered-storage-config", "", "Updates the configuration for managed tier checkpointing on the HyperPod cluster.")
		sagemaker_updateClusterCmd.MarkFlagRequired("cluster-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateClusterCmd)
}
