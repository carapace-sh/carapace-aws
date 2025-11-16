package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_batchDescribeModelPackageCmd = &cobra.Command{
	Use:   "batch-describe-model-package",
	Short: "This action batch describes a list of versioned model packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_batchDescribeModelPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_batchDescribeModelPackageCmd).Standalone()

		sagemaker_batchDescribeModelPackageCmd.Flags().String("model-package-arn-list", "", "The list of Amazon Resource Name (ARN) of the model package groups.")
		sagemaker_batchDescribeModelPackageCmd.MarkFlagRequired("model-package-arn-list")
	})
	sagemakerCmd.AddCommand(sagemaker_batchDescribeModelPackageCmd)
}
