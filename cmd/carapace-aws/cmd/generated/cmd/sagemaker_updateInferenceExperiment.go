package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateInferenceExperimentCmd = &cobra.Command{
	Use:   "update-inference-experiment",
	Short: "Updates an inference experiment that you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateInferenceExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateInferenceExperimentCmd).Standalone()

		sagemaker_updateInferenceExperimentCmd.Flags().String("data-storage-config", "", "The Amazon S3 location and configuration for storing inference request and response data.")
		sagemaker_updateInferenceExperimentCmd.Flags().String("description", "", "The description of the inference experiment.")
		sagemaker_updateInferenceExperimentCmd.Flags().String("model-variants", "", "An array of `ModelVariantConfig` objects.")
		sagemaker_updateInferenceExperimentCmd.Flags().String("name", "", "The name of the inference experiment to be updated.")
		sagemaker_updateInferenceExperimentCmd.Flags().String("schedule", "", "The duration for which the inference experiment will run.")
		sagemaker_updateInferenceExperimentCmd.Flags().String("shadow-mode-config", "", "The configuration of `ShadowMode` inference experiment type.")
		sagemaker_updateInferenceExperimentCmd.MarkFlagRequired("name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateInferenceExperimentCmd)
}
