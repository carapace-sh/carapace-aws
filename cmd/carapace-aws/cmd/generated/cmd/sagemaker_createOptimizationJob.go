package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createOptimizationJobCmd = &cobra.Command{
	Use:   "create-optimization-job",
	Short: "Creates a job that optimizes a model for inference performance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createOptimizationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createOptimizationJobCmd).Standalone()

		sagemaker_createOptimizationJobCmd.Flags().String("deployment-instance-type", "", "The type of instance that hosts the optimized model that you create with the optimization job.")
		sagemaker_createOptimizationJobCmd.Flags().String("model-source", "", "The location of the source model to optimize with an optimization job.")
		sagemaker_createOptimizationJobCmd.Flags().String("optimization-configs", "", "Settings for each of the optimization techniques that the job applies.")
		sagemaker_createOptimizationJobCmd.Flags().String("optimization-environment", "", "The environment variables to set in the model container.")
		sagemaker_createOptimizationJobCmd.Flags().String("optimization-job-name", "", "A custom name for the new optimization job.")
		sagemaker_createOptimizationJobCmd.Flags().String("output-config", "", "Details for where to store the optimized model that you create with the optimization job.")
		sagemaker_createOptimizationJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker AI to perform tasks on your behalf.")
		sagemaker_createOptimizationJobCmd.Flags().String("stopping-condition", "", "")
		sagemaker_createOptimizationJobCmd.Flags().String("tags", "", "A list of key-value pairs associated with the optimization job.")
		sagemaker_createOptimizationJobCmd.Flags().String("vpc-config", "", "A VPC in Amazon VPC that your optimized model has access to.")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("deployment-instance-type")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("model-source")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("optimization-configs")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("optimization-job-name")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("output-config")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("role-arn")
		sagemaker_createOptimizationJobCmd.MarkFlagRequired("stopping-condition")
	})
	sagemakerCmd.AddCommand(sagemaker_createOptimizationJobCmd)
}
