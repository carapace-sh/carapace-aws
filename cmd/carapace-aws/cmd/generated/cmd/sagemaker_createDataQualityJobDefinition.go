package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createDataQualityJobDefinitionCmd = &cobra.Command{
	Use:   "create-data-quality-job-definition",
	Short: "Creates a definition for a job that monitors data quality and drift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createDataQualityJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createDataQualityJobDefinitionCmd).Standalone()

		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("data-quality-app-specification", "", "Specifies the container that runs the monitoring job.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("data-quality-baseline-config", "", "Configures the constraints and baselines for the monitoring job.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("data-quality-job-input", "", "A list of inputs for the monitoring job.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("data-quality-job-output-config", "", "")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name for the monitoring job definition.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("job-resources", "", "")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("network-config", "", "Specifies networking configuration for the monitoring job.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform tasks on your behalf.")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("stopping-condition", "", "")
		sagemaker_createDataQualityJobDefinitionCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("data-quality-app-specification")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("data-quality-job-input")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("data-quality-job-output-config")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("job-resources")
		sagemaker_createDataQualityJobDefinitionCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createDataQualityJobDefinitionCmd)
}
