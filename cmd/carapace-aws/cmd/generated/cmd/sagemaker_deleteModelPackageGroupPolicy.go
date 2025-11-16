package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelPackageGroupPolicyCmd = &cobra.Command{
	Use:   "delete-model-package-group-policy",
	Short: "Deletes a model group resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelPackageGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelPackageGroupPolicyCmd).Standalone()

		sagemaker_deleteModelPackageGroupPolicyCmd.Flags().String("model-package-group-name", "", "The name of the model group for which to delete the policy.")
		sagemaker_deleteModelPackageGroupPolicyCmd.MarkFlagRequired("model-package-group-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelPackageGroupPolicyCmd)
}
