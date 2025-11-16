package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeClusterNodeCmd = &cobra.Command{
	Use:   "describe-cluster-node",
	Short: "Retrieves information of a node (also called a *instance* interchangeably) of a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeClusterNodeCmd).Standalone()

	sagemaker_describeClusterNodeCmd.Flags().String("cluster-name", "", "The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod cluster in which the node is.")
	sagemaker_describeClusterNodeCmd.Flags().String("node-id", "", "The ID of the SageMaker HyperPod cluster node.")
	sagemaker_describeClusterNodeCmd.Flags().String("node-logical-id", "", "The logical identifier of the node to describe.")
	sagemaker_describeClusterNodeCmd.MarkFlagRequired("cluster-name")
	sagemakerCmd.AddCommand(sagemaker_describeClusterNodeCmd)
}
