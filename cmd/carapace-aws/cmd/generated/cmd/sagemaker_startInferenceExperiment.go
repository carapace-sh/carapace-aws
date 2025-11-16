package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startInferenceExperimentCmd = &cobra.Command{
	Use:   "start-inference-experiment",
	Short: "Starts an inference experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startInferenceExperimentCmd).Standalone()

	sagemaker_startInferenceExperimentCmd.Flags().String("name", "", "The name of the inference experiment to start.")
	sagemaker_startInferenceExperimentCmd.MarkFlagRequired("name")
	sagemakerCmd.AddCommand(sagemaker_startInferenceExperimentCmd)
}
