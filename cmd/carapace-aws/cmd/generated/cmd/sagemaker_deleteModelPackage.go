package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelPackageCmd = &cobra.Command{
	Use:   "delete-model-package",
	Short: "Deletes a model package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelPackageCmd).Standalone()

		sagemaker_deleteModelPackageCmd.Flags().String("model-package-name", "", "The name or Amazon Resource Name (ARN) of the model package to delete.")
		sagemaker_deleteModelPackageCmd.MarkFlagRequired("model-package-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelPackageCmd)
}
