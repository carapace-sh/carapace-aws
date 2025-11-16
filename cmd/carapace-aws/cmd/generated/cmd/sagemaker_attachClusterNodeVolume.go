package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_attachClusterNodeVolumeCmd = &cobra.Command{
	Use:   "attach-cluster-node-volume",
	Short: "Attaches your Amazon Elastic Block Store (Amazon EBS) volume to a node in your EKS orchestrated HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_attachClusterNodeVolumeCmd).Standalone()

	sagemaker_attachClusterNodeVolumeCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of your SageMaker HyperPod cluster containing the target node.")
	sagemaker_attachClusterNodeVolumeCmd.Flags().String("node-id", "", "The unique identifier of the cluster node to which you want to attach the volume.")
	sagemaker_attachClusterNodeVolumeCmd.Flags().String("volume-id", "", "The unique identifier of your EBS volume to attach.")
	sagemaker_attachClusterNodeVolumeCmd.MarkFlagRequired("cluster-arn")
	sagemaker_attachClusterNodeVolumeCmd.MarkFlagRequired("node-id")
	sagemaker_attachClusterNodeVolumeCmd.MarkFlagRequired("volume-id")
	sagemakerCmd.AddCommand(sagemaker_attachClusterNodeVolumeCmd)
}
