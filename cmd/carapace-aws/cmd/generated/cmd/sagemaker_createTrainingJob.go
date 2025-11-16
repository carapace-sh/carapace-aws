package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createTrainingJobCmd = &cobra.Command{
	Use:   "create-training-job",
	Short: "Starts a model training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createTrainingJobCmd).Standalone()

	sagemaker_createTrainingJobCmd.Flags().String("algorithm-specification", "", "The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode.")
	sagemaker_createTrainingJobCmd.Flags().String("checkpoint-config", "", "Contains information about the output location for managed spot training checkpoint data.")
	sagemaker_createTrainingJobCmd.Flags().String("debug-hook-config", "", "")
	sagemaker_createTrainingJobCmd.Flags().String("debug-rule-configurations", "", "Configuration information for Amazon SageMaker Debugger rules for debugging output tensors.")
	sagemaker_createTrainingJobCmd.Flags().Bool("enable-inter-container-traffic-encryption", false, "To encrypt all communications between ML compute instances in distributed training, choose `True`.")
	sagemaker_createTrainingJobCmd.Flags().Bool("enable-managed-spot-training", false, "To train models using managed spot training, choose `True`.")
	sagemaker_createTrainingJobCmd.Flags().Bool("enable-network-isolation", false, "Isolates the training container.")
	sagemaker_createTrainingJobCmd.Flags().String("environment", "", "The environment variables to set in the Docker container.")
	sagemaker_createTrainingJobCmd.Flags().String("experiment-config", "", "")
	sagemaker_createTrainingJobCmd.Flags().String("hyper-parameters", "", "Algorithm-specific parameters that influence the quality of the model.")
	sagemaker_createTrainingJobCmd.Flags().String("infra-check-config", "", "Contains information about the infrastructure health check configuration for the training job.")
	sagemaker_createTrainingJobCmd.Flags().String("input-data-config", "", "An array of `Channel` objects.")
	sagemaker_createTrainingJobCmd.Flags().Bool("no-enable-inter-container-traffic-encryption", false, "To encrypt all communications between ML compute instances in distributed training, choose `True`.")
	sagemaker_createTrainingJobCmd.Flags().Bool("no-enable-managed-spot-training", false, "To train models using managed spot training, choose `True`.")
	sagemaker_createTrainingJobCmd.Flags().Bool("no-enable-network-isolation", false, "Isolates the training container.")
	sagemaker_createTrainingJobCmd.Flags().String("output-data-config", "", "Specifies the path to the S3 location where you want to store model artifacts.")
	sagemaker_createTrainingJobCmd.Flags().String("profiler-config", "", "")
	sagemaker_createTrainingJobCmd.Flags().String("profiler-rule-configurations", "", "Configuration information for Amazon SageMaker Debugger rules for profiling system and framework metrics.")
	sagemaker_createTrainingJobCmd.Flags().String("remote-debug-config", "", "Configuration for remote debugging.")
	sagemaker_createTrainingJobCmd.Flags().String("resource-config", "", "The resources, including the ML compute instances and ML storage volumes, to use for model training.")
	sagemaker_createTrainingJobCmd.Flags().String("retry-strategy", "", "The number of times to retry the job when the job fails due to an `InternalServerError`.")
	sagemaker_createTrainingJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.")
	sagemaker_createTrainingJobCmd.Flags().String("session-chaining-config", "", "Contains information about attribute-based access control (ABAC) for the training job.")
	sagemaker_createTrainingJobCmd.Flags().String("stopping-condition", "", "Specifies a limit to how long a model training job can run.")
	sagemaker_createTrainingJobCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createTrainingJobCmd.Flags().String("tensor-board-output-config", "", "")
	sagemaker_createTrainingJobCmd.Flags().String("training-job-name", "", "The name of the training job.")
	sagemaker_createTrainingJobCmd.Flags().String("vpc-config", "", "A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html) object that specifies the VPC that you want your training job to connect to.")
	sagemaker_createTrainingJobCmd.Flag("no-enable-inter-container-traffic-encryption").Hidden = true
	sagemaker_createTrainingJobCmd.Flag("no-enable-managed-spot-training").Hidden = true
	sagemaker_createTrainingJobCmd.Flag("no-enable-network-isolation").Hidden = true
	sagemaker_createTrainingJobCmd.MarkFlagRequired("output-data-config")
	sagemaker_createTrainingJobCmd.MarkFlagRequired("role-arn")
	sagemaker_createTrainingJobCmd.MarkFlagRequired("training-job-name")
	sagemakerCmd.AddCommand(sagemaker_createTrainingJobCmd)
}
