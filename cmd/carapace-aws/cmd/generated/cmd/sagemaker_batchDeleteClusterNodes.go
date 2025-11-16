package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_batchDeleteClusterNodesCmd = &cobra.Command{
	Use:   "batch-delete-cluster-nodes",
	Short: "Deletes specific nodes within a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_batchDeleteClusterNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_batchDeleteClusterNodesCmd).Standalone()

		sagemaker_batchDeleteClusterNodesCmd.Flags().String("cluster-name", "", "The name of the SageMaker HyperPod cluster from which to delete the specified nodes.")
		sagemaker_batchDeleteClusterNodesCmd.Flags().String("node-ids", "", "A list of node IDs to be deleted from the specified cluster.")
		sagemaker_batchDeleteClusterNodesCmd.Flags().String("node-logical-ids", "", "A list of `NodeLogicalIds` identifying the nodes to be deleted.")
		sagemaker_batchDeleteClusterNodesCmd.MarkFlagRequired("cluster-name")
	})
	sagemakerCmd.AddCommand(sagemaker_batchDeleteClusterNodesCmd)
}
