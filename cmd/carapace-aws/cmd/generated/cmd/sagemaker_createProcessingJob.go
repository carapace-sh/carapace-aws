package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createProcessingJobCmd = &cobra.Command{
	Use:   "create-processing-job",
	Short: "Creates a processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createProcessingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createProcessingJobCmd).Standalone()

		sagemaker_createProcessingJobCmd.Flags().String("app-specification", "", "Configures the processing job to run a specified Docker container image.")
		sagemaker_createProcessingJobCmd.Flags().String("environment", "", "The environment variables to set in the Docker container.")
		sagemaker_createProcessingJobCmd.Flags().String("experiment-config", "", "")
		sagemaker_createProcessingJobCmd.Flags().String("network-config", "", "Networking options for a processing job, such as whether to allow inbound and outbound network calls to and from processing containers, and the VPC subnets and security groups to use for VPC-enabled processing jobs.")
		sagemaker_createProcessingJobCmd.Flags().String("processing-inputs", "", "An array of inputs configuring the data to download into the processing container.")
		sagemaker_createProcessingJobCmd.Flags().String("processing-job-name", "", "The name of the processing job.")
		sagemaker_createProcessingJobCmd.Flags().String("processing-output-config", "", "Output configuration for the processing job.")
		sagemaker_createProcessingJobCmd.Flags().String("processing-resources", "", "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.")
		sagemaker_createProcessingJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.")
		sagemaker_createProcessingJobCmd.Flags().String("stopping-condition", "", "The time limit for how long the processing job is allowed to run.")
		sagemaker_createProcessingJobCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
		sagemaker_createProcessingJobCmd.MarkFlagRequired("app-specification")
		sagemaker_createProcessingJobCmd.MarkFlagRequired("processing-job-name")
		sagemaker_createProcessingJobCmd.MarkFlagRequired("processing-resources")
		sagemaker_createProcessingJobCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createProcessingJobCmd)
}
