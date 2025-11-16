package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopInferenceExperimentCmd = &cobra.Command{
	Use:   "stop-inference-experiment",
	Short: "Stops an inference experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopInferenceExperimentCmd).Standalone()

	sagemaker_stopInferenceExperimentCmd.Flags().String("desired-model-variants", "", "An array of `ModelVariantConfig` objects.")
	sagemaker_stopInferenceExperimentCmd.Flags().String("desired-state", "", "The desired state of the experiment after stopping.")
	sagemaker_stopInferenceExperimentCmd.Flags().String("model-variant-actions", "", "Array of key-value pairs, with names of variants mapped to actions.")
	sagemaker_stopInferenceExperimentCmd.Flags().String("name", "", "The name of the inference experiment to stop.")
	sagemaker_stopInferenceExperimentCmd.Flags().String("reason", "", "The reason for stopping the experiment.")
	sagemaker_stopInferenceExperimentCmd.MarkFlagRequired("model-variant-actions")
	sagemaker_stopInferenceExperimentCmd.MarkFlagRequired("name")
	sagemakerCmd.AddCommand(sagemaker_stopInferenceExperimentCmd)
}
