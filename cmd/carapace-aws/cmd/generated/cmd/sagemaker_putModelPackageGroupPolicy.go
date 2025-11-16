package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_putModelPackageGroupPolicyCmd = &cobra.Command{
	Use:   "put-model-package-group-policy",
	Short: "Adds a resouce policy to control access to a model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_putModelPackageGroupPolicyCmd).Standalone()

	sagemaker_putModelPackageGroupPolicyCmd.Flags().String("model-package-group-name", "", "The name of the model group to add a resource policy to.")
	sagemaker_putModelPackageGroupPolicyCmd.Flags().String("resource-policy", "", "The resource policy for the model group.")
	sagemaker_putModelPackageGroupPolicyCmd.MarkFlagRequired("model-package-group-name")
	sagemaker_putModelPackageGroupPolicyCmd.MarkFlagRequired("resource-policy")
	sagemakerCmd.AddCommand(sagemaker_putModelPackageGroupPolicyCmd)
}
