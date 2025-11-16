package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Creates a model in SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createModelCmd).Standalone()

		sagemaker_createModelCmd.Flags().String("containers", "", "Specifies the containers in the inference pipeline.")
		sagemaker_createModelCmd.Flags().Bool("enable-network-isolation", false, "Isolates the model container.")
		sagemaker_createModelCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs.")
		sagemaker_createModelCmd.Flags().String("inference-execution-config", "", "Specifies details of how containers in a multi-container endpoint are called.")
		sagemaker_createModelCmd.Flags().String("model-name", "", "The name of the new model.")
		sagemaker_createModelCmd.Flags().Bool("no-enable-network-isolation", false, "Isolates the model container.")
		sagemaker_createModelCmd.Flags().String("primary-container", "", "The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.")
		sagemaker_createModelCmd.Flags().String("tags", "", "An array of key-value pairs.")
		sagemaker_createModelCmd.Flags().String("vpc-config", "", "A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html) object that specifies the VPC that you want your model to connect to.")
		sagemaker_createModelCmd.MarkFlagRequired("model-name")
		sagemaker_createModelCmd.Flag("no-enable-network-isolation").Hidden = true
	})
	sagemakerCmd.AddCommand(sagemaker_createModelCmd)
}
