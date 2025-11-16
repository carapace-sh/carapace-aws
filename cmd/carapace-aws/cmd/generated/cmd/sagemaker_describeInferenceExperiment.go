package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeInferenceExperimentCmd = &cobra.Command{
	Use:   "describe-inference-experiment",
	Short: "Returns details about an inference experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeInferenceExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeInferenceExperimentCmd).Standalone()

		sagemaker_describeInferenceExperimentCmd.Flags().String("name", "", "The name of the inference experiment to describe.")
		sagemaker_describeInferenceExperimentCmd.MarkFlagRequired("name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeInferenceExperimentCmd)
}
