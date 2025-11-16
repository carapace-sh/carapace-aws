package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelBiasJobDefinitionCmd = &cobra.Command{
	Use:   "create-model-bias-job-definition",
	Short: "Creates the definition for a model bias job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelBiasJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createModelBiasJobDefinitionCmd).Standalone()

		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the bias job definition.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("job-resources", "", "")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("model-bias-app-specification", "", "Configures the model bias job to run a specified Docker container image.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("model-bias-baseline-config", "", "The baseline configuration for a model bias job.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("model-bias-job-input", "", "Inputs for the model bias job.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("model-bias-job-output-config", "", "")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("network-config", "", "Networking options for a model bias job.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform tasks on your behalf.")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("stopping-condition", "", "")
		sagemaker_createModelBiasJobDefinitionCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("job-definition-name")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("job-resources")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("model-bias-app-specification")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("model-bias-job-input")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("model-bias-job-output-config")
		sagemaker_createModelBiasJobDefinitionCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createModelBiasJobDefinitionCmd)
}
