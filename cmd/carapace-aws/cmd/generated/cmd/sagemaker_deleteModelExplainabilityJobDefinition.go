package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelExplainabilityJobDefinitionCmd = &cobra.Command{
	Use:   "delete-model-explainability-job-definition",
	Short: "Deletes an Amazon SageMaker AI model explainability job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelExplainabilityJobDefinitionCmd).Standalone()

	sagemaker_deleteModelExplainabilityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model explainability job definition to delete.")
	sagemaker_deleteModelExplainabilityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemakerCmd.AddCommand(sagemaker_deleteModelExplainabilityJobDefinitionCmd)
}
