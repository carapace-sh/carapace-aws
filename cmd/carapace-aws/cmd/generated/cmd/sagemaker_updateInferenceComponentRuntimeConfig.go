package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateInferenceComponentRuntimeConfigCmd = &cobra.Command{
	Use:   "update-inference-component-runtime-config",
	Short: "Runtime settings for a model that is deployed with an inference component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateInferenceComponentRuntimeConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateInferenceComponentRuntimeConfigCmd).Standalone()

		sagemaker_updateInferenceComponentRuntimeConfigCmd.Flags().String("desired-runtime-config", "", "Runtime settings for a model that is deployed with an inference component.")
		sagemaker_updateInferenceComponentRuntimeConfigCmd.Flags().String("inference-component-name", "", "The name of the inference component to update.")
		sagemaker_updateInferenceComponentRuntimeConfigCmd.MarkFlagRequired("desired-runtime-config")
		sagemaker_updateInferenceComponentRuntimeConfigCmd.MarkFlagRequired("inference-component-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateInferenceComponentRuntimeConfigCmd)
}
