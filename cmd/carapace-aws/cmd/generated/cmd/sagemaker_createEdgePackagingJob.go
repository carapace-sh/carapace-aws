package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createEdgePackagingJobCmd = &cobra.Command{
	Use:   "create-edge-packaging-job",
	Short: "Starts a SageMaker Edge Manager model packaging job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createEdgePackagingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createEdgePackagingJobCmd).Standalone()

		sagemaker_createEdgePackagingJobCmd.Flags().String("compilation-job-name", "", "The name of the SageMaker Neo compilation job that will be used to locate model artifacts for packaging.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("edge-packaging-job-name", "", "The name of the edge packaging job.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("model-name", "", "The name of the model.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("model-version", "", "The version of the model.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("output-config", "", "Provides information about the output location for the packaged model.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("resource-key", "", "The Amazon Web Services KMS key to use when encrypting the EBS volume the edge packaging job runs on.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to download and upload the model, and to contact SageMaker Neo.")
		sagemaker_createEdgePackagingJobCmd.Flags().String("tags", "", "Creates tags for the packaging job.")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("compilation-job-name")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("edge-packaging-job-name")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("model-name")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("model-version")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("output-config")
		sagemaker_createEdgePackagingJobCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createEdgePackagingJobCmd)
}
