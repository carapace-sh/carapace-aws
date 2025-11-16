package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createInferenceExperimentCmd = &cobra.Command{
	Use:   "create-inference-experiment",
	Short: "Creates an inference experiment using the configurations specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createInferenceExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createInferenceExperimentCmd).Standalone()

		sagemaker_createInferenceExperimentCmd.Flags().String("data-storage-config", "", "The Amazon S3 location and configuration for storing inference request and response data.")
		sagemaker_createInferenceExperimentCmd.Flags().String("description", "", "A description for the inference experiment.")
		sagemaker_createInferenceExperimentCmd.Flags().String("endpoint-name", "", "The name of the Amazon SageMaker endpoint on which you want to run the inference experiment.")
		sagemaker_createInferenceExperimentCmd.Flags().String("kms-key", "", "The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.")
		sagemaker_createInferenceExperimentCmd.Flags().String("model-variants", "", "An array of `ModelVariantConfig` objects.")
		sagemaker_createInferenceExperimentCmd.Flags().String("name", "", "The name for the inference experiment.")
		sagemaker_createInferenceExperimentCmd.Flags().String("role-arn", "", "The ARN of the IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.")
		sagemaker_createInferenceExperimentCmd.Flags().String("schedule", "", "The duration for which you want the inference experiment to run.")
		sagemaker_createInferenceExperimentCmd.Flags().String("shadow-mode-config", "", "The configuration of `ShadowMode` inference experiment type.")
		sagemaker_createInferenceExperimentCmd.Flags().String("tags", "", "Array of key-value pairs.")
		sagemaker_createInferenceExperimentCmd.Flags().String("type", "", "The type of the inference experiment that you want to run.")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("endpoint-name")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("model-variants")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("name")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("role-arn")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("shadow-mode-config")
		sagemaker_createInferenceExperimentCmd.MarkFlagRequired("type")
	})
	sagemakerCmd.AddCommand(sagemaker_createInferenceExperimentCmd)
}
