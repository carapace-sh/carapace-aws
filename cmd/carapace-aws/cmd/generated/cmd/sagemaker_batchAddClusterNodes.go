package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_batchAddClusterNodesCmd = &cobra.Command{
	Use:   "batch-add-cluster-nodes",
	Short: "Adds nodes to a HyperPod cluster by incrementing the target count for one or more instance groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_batchAddClusterNodesCmd).Standalone()

	sagemaker_batchAddClusterNodesCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	sagemaker_batchAddClusterNodesCmd.Flags().String("cluster-name", "", "The name of the HyperPod cluster to which you want to add nodes.")
	sagemaker_batchAddClusterNodesCmd.Flags().String("nodes-to-add", "", "A list of instance groups and the number of nodes to add to each.")
	sagemaker_batchAddClusterNodesCmd.MarkFlagRequired("cluster-name")
	sagemaker_batchAddClusterNodesCmd.MarkFlagRequired("nodes-to-add")
	sagemakerCmd.AddCommand(sagemaker_batchAddClusterNodesCmd)
}
