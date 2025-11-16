package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelPackageGroupCmd = &cobra.Command{
	Use:   "delete-model-package-group",
	Short: "Deletes the specified model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelPackageGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelPackageGroupCmd).Standalone()

		sagemaker_deleteModelPackageGroupCmd.Flags().String("model-package-group-name", "", "The name of the model group to delete.")
		sagemaker_deleteModelPackageGroupCmd.MarkFlagRequired("model-package-group-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelPackageGroupCmd)
}
