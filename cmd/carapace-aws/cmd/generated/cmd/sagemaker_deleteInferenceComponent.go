package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteInferenceComponentCmd = &cobra.Command{
	Use:   "delete-inference-component",
	Short: "Deletes an inference component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteInferenceComponentCmd).Standalone()

	sagemaker_deleteInferenceComponentCmd.Flags().String("inference-component-name", "", "The name of the inference component to delete.")
	sagemaker_deleteInferenceComponentCmd.MarkFlagRequired("inference-component-name")
	sagemakerCmd.AddCommand(sagemaker_deleteInferenceComponentCmd)
}
