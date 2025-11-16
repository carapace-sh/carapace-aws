package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createStudioLifecycleConfigCmd = &cobra.Command{
	Use:   "create-studio-lifecycle-config",
	Short: "Creates a new Amazon SageMaker AI Studio Lifecycle Configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createStudioLifecycleConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createStudioLifecycleConfigCmd).Standalone()

		sagemaker_createStudioLifecycleConfigCmd.Flags().String("studio-lifecycle-config-app-type", "", "The App type that the Lifecycle Configuration is attached to.")
		sagemaker_createStudioLifecycleConfigCmd.Flags().String("studio-lifecycle-config-content", "", "The content of your Amazon SageMaker AI Studio Lifecycle Configuration script.")
		sagemaker_createStudioLifecycleConfigCmd.Flags().String("studio-lifecycle-config-name", "", "The name of the Amazon SageMaker AI Studio Lifecycle Configuration to create.")
		sagemaker_createStudioLifecycleConfigCmd.Flags().String("tags", "", "Tags to be associated with the Lifecycle Configuration.")
		sagemaker_createStudioLifecycleConfigCmd.MarkFlagRequired("studio-lifecycle-config-app-type")
		sagemaker_createStudioLifecycleConfigCmd.MarkFlagRequired("studio-lifecycle-config-content")
		sagemaker_createStudioLifecycleConfigCmd.MarkFlagRequired("studio-lifecycle-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createStudioLifecycleConfigCmd)
}
