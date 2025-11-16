package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelPackageGroupCmd = &cobra.Command{
	Use:   "describe-model-package-group",
	Short: "Gets a description for the specified model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelPackageGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeModelPackageGroupCmd).Standalone()

		sagemaker_describeModelPackageGroupCmd.Flags().String("model-package-group-name", "", "The name of the model group to describe.")
		sagemaker_describeModelPackageGroupCmd.MarkFlagRequired("model-package-group-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeModelPackageGroupCmd)
}
