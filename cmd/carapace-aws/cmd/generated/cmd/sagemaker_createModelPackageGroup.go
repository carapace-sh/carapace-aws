package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelPackageGroupCmd = &cobra.Command{
	Use:   "create-model-package-group",
	Short: "Creates a model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelPackageGroupCmd).Standalone()

	sagemaker_createModelPackageGroupCmd.Flags().String("model-package-group-description", "", "A description for the model group.")
	sagemaker_createModelPackageGroupCmd.Flags().String("model-package-group-name", "", "The name of the model group.")
	sagemaker_createModelPackageGroupCmd.Flags().String("tags", "", "A list of key value pairs associated with the model group.")
	sagemaker_createModelPackageGroupCmd.MarkFlagRequired("model-package-group-name")
	sagemakerCmd.AddCommand(sagemaker_createModelPackageGroupCmd)
}
