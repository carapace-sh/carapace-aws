package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createClusterCmd).Standalone()

		sagemaker_createClusterCmd.Flags().String("auto-scaling", "", "The autoscaling configuration for the cluster.")
		sagemaker_createClusterCmd.Flags().String("cluster-name", "", "The name for the new SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("cluster-role", "", "The Amazon Resource Name (ARN) of the IAM role that HyperPod assumes to perform cluster autoscaling operations.")
		sagemaker_createClusterCmd.Flags().String("instance-groups", "", "The instance groups to be created in the SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("node-provisioning-mode", "", "The mode for provisioning nodes in the cluster.")
		sagemaker_createClusterCmd.Flags().String("node-recovery", "", "The node recovery mode for the SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("orchestrator", "", "The type of orchestrator to use for the SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("restricted-instance-groups", "", "The specialized instance groups for training models like Amazon Nova to be created in the SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("tags", "", "Custom tags for managing the SageMaker HyperPod cluster as an Amazon Web Services resource.")
		sagemaker_createClusterCmd.Flags().String("tiered-storage-config", "", "The configuration for managed tier checkpointing on the HyperPod cluster.")
		sagemaker_createClusterCmd.Flags().String("vpc-config", "", "Specifies the Amazon Virtual Private Cloud (VPC) that is associated with the Amazon SageMaker HyperPod cluster.")
		sagemaker_createClusterCmd.MarkFlagRequired("cluster-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createClusterCmd)
}
