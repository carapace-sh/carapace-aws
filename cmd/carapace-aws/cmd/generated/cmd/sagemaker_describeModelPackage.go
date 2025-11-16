package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelPackageCmd = &cobra.Command{
	Use:   "describe-model-package",
	Short: "Returns a description of the specified model package, which is used to create SageMaker models or list them on Amazon Web Services Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelPackageCmd).Standalone()

	sagemaker_describeModelPackageCmd.Flags().String("model-package-name", "", "The name or Amazon Resource Name (ARN) of the model package to describe.")
	sagemaker_describeModelPackageCmd.MarkFlagRequired("model-package-name")
	sagemakerCmd.AddCommand(sagemaker_describeModelPackageCmd)
}
