package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getModelPackageGroupPolicyCmd = &cobra.Command{
	Use:   "get-model-package-group-policy",
	Short: "Gets a resource policy that manages access for a model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getModelPackageGroupPolicyCmd).Standalone()

	sagemaker_getModelPackageGroupPolicyCmd.Flags().String("model-package-group-name", "", "The name of the model group for which to get the resource policy.")
	sagemaker_getModelPackageGroupPolicyCmd.MarkFlagRequired("model-package-group-name")
	sagemakerCmd.AddCommand(sagemaker_getModelPackageGroupPolicyCmd)
}
