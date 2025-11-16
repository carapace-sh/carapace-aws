package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeInferenceComponentCmd = &cobra.Command{
	Use:   "describe-inference-component",
	Short: "Returns information about an inference component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeInferenceComponentCmd).Standalone()

	sagemaker_describeInferenceComponentCmd.Flags().String("inference-component-name", "", "The name of the inference component.")
	sagemaker_describeInferenceComponentCmd.MarkFlagRequired("inference-component-name")
	sagemakerCmd.AddCommand(sagemaker_describeInferenceComponentCmd)
}
