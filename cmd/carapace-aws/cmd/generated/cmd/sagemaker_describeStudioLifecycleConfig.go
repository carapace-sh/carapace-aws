package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeStudioLifecycleConfigCmd = &cobra.Command{
	Use:   "describe-studio-lifecycle-config",
	Short: "Describes the Amazon SageMaker AI Studio Lifecycle Configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeStudioLifecycleConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeStudioLifecycleConfigCmd).Standalone()

		sagemaker_describeStudioLifecycleConfigCmd.Flags().String("studio-lifecycle-config-name", "", "The name of the Amazon SageMaker AI Studio Lifecycle Configuration to describe.")
		sagemaker_describeStudioLifecycleConfigCmd.MarkFlagRequired("studio-lifecycle-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeStudioLifecycleConfigCmd)
}
