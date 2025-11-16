package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateInferenceComponentCmd = &cobra.Command{
	Use:   "update-inference-component",
	Short: "Updates an inference component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateInferenceComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateInferenceComponentCmd).Standalone()

		sagemaker_updateInferenceComponentCmd.Flags().String("deployment-config", "", "The deployment configuration for the inference component.")
		sagemaker_updateInferenceComponentCmd.Flags().String("inference-component-name", "", "The name of the inference component.")
		sagemaker_updateInferenceComponentCmd.Flags().String("runtime-config", "", "Runtime settings for a model that is deployed with an inference component.")
		sagemaker_updateInferenceComponentCmd.Flags().String("specification", "", "Details about the resources to deploy with this inference component, including the model, container, and compute resources.")
		sagemaker_updateInferenceComponentCmd.MarkFlagRequired("inference-component-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateInferenceComponentCmd)
}
