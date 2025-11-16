package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteStudioLifecycleConfigCmd = &cobra.Command{
	Use:   "delete-studio-lifecycle-config",
	Short: "Deletes the Amazon SageMaker AI Studio Lifecycle Configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteStudioLifecycleConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteStudioLifecycleConfigCmd).Standalone()

		sagemaker_deleteStudioLifecycleConfigCmd.Flags().String("studio-lifecycle-config-name", "", "The name of the Amazon SageMaker AI Studio Lifecycle Configuration to delete.")
		sagemaker_deleteStudioLifecycleConfigCmd.MarkFlagRequired("studio-lifecycle-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteStudioLifecycleConfigCmd)
}
