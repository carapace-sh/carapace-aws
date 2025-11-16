package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteInferenceExperimentCmd = &cobra.Command{
	Use:   "delete-inference-experiment",
	Short: "Deletes an inference experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteInferenceExperimentCmd).Standalone()

	sagemaker_deleteInferenceExperimentCmd.Flags().String("name", "", "The name of the inference experiment you want to delete.")
	sagemaker_deleteInferenceExperimentCmd.MarkFlagRequired("name")
	sagemakerCmd.AddCommand(sagemaker_deleteInferenceExperimentCmd)
}
