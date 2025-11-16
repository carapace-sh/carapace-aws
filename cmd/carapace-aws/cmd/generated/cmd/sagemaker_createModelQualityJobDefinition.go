package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelQualityJobDefinitionCmd = &cobra.Command{
	Use:   "create-model-quality-job-definition",
	Short: "Creates a definition for a job that monitors model quality and drift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelQualityJobDefinitionCmd).Standalone()

	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the monitoring job definition.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("job-resources", "", "")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("model-quality-app-specification", "", "The container that runs the monitoring job.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("model-quality-baseline-config", "", "Specifies the constraints and baselines for the monitoring job.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("model-quality-job-input", "", "A list of the inputs that are monitored.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("model-quality-job-output-config", "", "")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("network-config", "", "Specifies the network configuration for the monitoring job.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform tasks on your behalf.")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("stopping-condition", "", "")
	sagemaker_createModelQualityJobDefinitionCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("job-resources")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("model-quality-app-specification")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("model-quality-job-input")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("model-quality-job-output-config")
	sagemaker_createModelQualityJobDefinitionCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createModelQualityJobDefinitionCmd)
}
