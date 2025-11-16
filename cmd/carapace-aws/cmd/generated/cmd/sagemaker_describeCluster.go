package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Retrieves information of a SageMaker HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeClusterCmd).Standalone()

		sagemaker_describeClusterCmd.Flags().String("cluster-name", "", "The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod cluster.")
		sagemaker_describeClusterCmd.MarkFlagRequired("cluster-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeClusterCmd)
}
