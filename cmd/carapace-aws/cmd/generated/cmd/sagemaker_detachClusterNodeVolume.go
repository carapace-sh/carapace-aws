package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_detachClusterNodeVolumeCmd = &cobra.Command{
	Use:   "detach-cluster-node-volume",
	Short: "Detaches your Amazon Elastic Block Store (Amazon EBS) volume from a node in your EKS orchestrated SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_detachClusterNodeVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_detachClusterNodeVolumeCmd).Standalone()

		sagemaker_detachClusterNodeVolumeCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of your SageMaker HyperPod cluster containing the target node.")
		sagemaker_detachClusterNodeVolumeCmd.Flags().String("node-id", "", "The unique identifier of the cluster node from which you want to detach the volume.")
		sagemaker_detachClusterNodeVolumeCmd.Flags().String("volume-id", "", "The unique identifier of your EBS volume that you want to detach.")
		sagemaker_detachClusterNodeVolumeCmd.MarkFlagRequired("cluster-arn")
		sagemaker_detachClusterNodeVolumeCmd.MarkFlagRequired("node-id")
		sagemaker_detachClusterNodeVolumeCmd.MarkFlagRequired("volume-id")
	})
	sagemakerCmd.AddCommand(sagemaker_detachClusterNodeVolumeCmd)
}
