package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateClusterSoftwareCmd = &cobra.Command{
	Use:   "update-cluster-software",
	Short: "Updates the platform software of a SageMaker HyperPod cluster for security patching.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateClusterSoftwareCmd).Standalone()

	sagemaker_updateClusterSoftwareCmd.Flags().String("cluster-name", "", "Specify the name or the Amazon Resource Name (ARN) of the SageMaker HyperPod cluster you want to update for security patching.")
	sagemaker_updateClusterSoftwareCmd.Flags().String("deployment-config", "", "The configuration to use when updating the AMI versions.")
	sagemaker_updateClusterSoftwareCmd.Flags().String("image-id", "", "When configuring your HyperPod cluster, you can specify an image ID using one of the following options:")
	sagemaker_updateClusterSoftwareCmd.Flags().String("instance-groups", "", "The array of instance groups for which to update AMI versions.")
	sagemaker_updateClusterSoftwareCmd.MarkFlagRequired("cluster-name")
	sagemakerCmd.AddCommand(sagemaker_updateClusterSoftwareCmd)
}
