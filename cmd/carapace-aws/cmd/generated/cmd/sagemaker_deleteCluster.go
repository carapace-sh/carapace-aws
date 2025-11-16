package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Delete a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteClusterCmd).Standalone()

		sagemaker_deleteClusterCmd.Flags().String("cluster-name", "", "The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod cluster to delete.")
		sagemaker_deleteClusterCmd.MarkFlagRequired("cluster-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteClusterCmd)
}
