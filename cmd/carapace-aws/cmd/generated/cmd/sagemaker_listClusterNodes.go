package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listClusterNodesCmd = &cobra.Command{
	Use:   "list-cluster-nodes",
	Short: "Retrieves the list of instances (also called *nodes* interchangeably) in a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listClusterNodesCmd).Standalone()

	sagemaker_listClusterNodesCmd.Flags().String("cluster-name", "", "The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod cluster in which you want to retrieve the list of nodes.")
	sagemaker_listClusterNodesCmd.Flags().String("creation-time-after", "", "A filter that returns nodes in a SageMaker HyperPod cluster created after the specified time.")
	sagemaker_listClusterNodesCmd.Flags().String("creation-time-before", "", "A filter that returns nodes in a SageMaker HyperPod cluster created before the specified time.")
	sagemaker_listClusterNodesCmd.Flags().String("include-node-logical-ids", "", "Specifies whether to include nodes that are still being provisioned in the response.")
	sagemaker_listClusterNodesCmd.Flags().String("instance-group-name-contains", "", "A filter that returns the instance groups whose name contain a specified string.")
	sagemaker_listClusterNodesCmd.Flags().String("max-results", "", "The maximum number of nodes to return in the response.")
	sagemaker_listClusterNodesCmd.Flags().String("next-token", "", "If the result of the previous `ListClusterNodes` request was truncated, the response includes a `NextToken`.")
	sagemaker_listClusterNodesCmd.Flags().String("sort-by", "", "The field by which to sort results.")
	sagemaker_listClusterNodesCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemaker_listClusterNodesCmd.MarkFlagRequired("cluster-name")
	sagemakerCmd.AddCommand(sagemaker_listClusterNodesCmd)
}
