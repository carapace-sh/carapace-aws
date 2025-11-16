package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelExplainabilityJobDefinitionCmd = &cobra.Command{
	Use:   "describe-model-explainability-job-definition",
	Short: "Returns a description of a model explainability job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelExplainabilityJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeModelExplainabilityJobDefinitionCmd).Standalone()

		sagemaker_describeModelExplainabilityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model explainability job definition.")
		sagemaker_describeModelExplainabilityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeModelExplainabilityJobDefinitionCmd)
}
