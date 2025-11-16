package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelExplainabilityJobDefinitionCmd = &cobra.Command{
	Use:   "create-model-explainability-job-definition",
	Short: "Creates the definition for a model explainability job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelExplainabilityJobDefinitionCmd).Standalone()

	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model explainability job definition.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("job-resources", "", "")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("model-explainability-app-specification", "", "Configures the model explainability job to run a specified Docker container image.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("model-explainability-baseline-config", "", "The baseline configuration for a model explainability job.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("model-explainability-job-input", "", "Inputs for the model explainability job.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("model-explainability-job-output-config", "", "")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("network-config", "", "Networking options for a model explainability job.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform tasks on your behalf.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("stopping-condition", "", "")
	sagemaker_createModelExplainabilityJobDefinitionCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("job-resources")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("model-explainability-app-specification")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("model-explainability-job-input")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("model-explainability-job-output-config")
	sagemaker_createModelExplainabilityJobDefinitionCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createModelExplainabilityJobDefinitionCmd)
}
