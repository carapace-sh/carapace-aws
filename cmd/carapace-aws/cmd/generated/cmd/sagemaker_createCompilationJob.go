package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createCompilationJobCmd = &cobra.Command{
	Use:   "create-compilation-job",
	Short: "Starts a model compilation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createCompilationJobCmd).Standalone()

	sagemaker_createCompilationJobCmd.Flags().String("compilation-job-name", "", "A name for the model compilation job.")
	sagemaker_createCompilationJobCmd.Flags().String("input-config", "", "Provides information about the location of input model artifacts, the name and shape of the expected data inputs, and the framework in which the model was trained.")
	sagemaker_createCompilationJobCmd.Flags().String("model-package-version-arn", "", "The Amazon Resource Name (ARN) of a versioned model package.")
	sagemaker_createCompilationJobCmd.Flags().String("output-config", "", "Provides information about the output location for the compiled model and the target device the model runs on.")
	sagemaker_createCompilationJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker AI to perform tasks on your behalf.")
	sagemaker_createCompilationJobCmd.Flags().String("stopping-condition", "", "Specifies a limit to how long a model compilation job can run.")
	sagemaker_createCompilationJobCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createCompilationJobCmd.Flags().String("vpc-config", "", "A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html) object that specifies the VPC that you want your compilation job to connect to.")
	sagemaker_createCompilationJobCmd.MarkFlagRequired("compilation-job-name")
	sagemaker_createCompilationJobCmd.MarkFlagRequired("output-config")
	sagemaker_createCompilationJobCmd.MarkFlagRequired("role-arn")
	sagemaker_createCompilationJobCmd.MarkFlagRequired("stopping-condition")
	sagemakerCmd.AddCommand(sagemaker_createCompilationJobCmd)
}
