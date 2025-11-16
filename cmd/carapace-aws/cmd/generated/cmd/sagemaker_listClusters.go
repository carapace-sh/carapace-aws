package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Retrieves the list of SageMaker HyperPod clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listClustersCmd).Standalone()

		sagemaker_listClustersCmd.Flags().String("creation-time-after", "", "Set a start time for the time range during which you want to list SageMaker HyperPod clusters.")
		sagemaker_listClustersCmd.Flags().String("creation-time-before", "", "Set an end time for the time range during which you want to list SageMaker HyperPod clusters.")
		sagemaker_listClustersCmd.Flags().String("max-results", "", "Specifies the maximum number of clusters to evaluate for the operation (not necessarily the number of matching items).")
		sagemaker_listClustersCmd.Flags().String("name-contains", "", "Set the maximum number of instances to print in the list.")
		sagemaker_listClustersCmd.Flags().String("next-token", "", "Set the next token to retrieve the list of SageMaker HyperPod clusters.")
		sagemaker_listClustersCmd.Flags().String("sort-by", "", "The field by which to sort results.")
		sagemaker_listClustersCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listClustersCmd.Flags().String("training-plan-arn", "", "The Amazon Resource Name (ARN); of the training plan to filter clusters by.")
	})
	sagemakerCmd.AddCommand(sagemaker_listClustersCmd)
}
