package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelPackageCmd = &cobra.Command{
	Use:   "create-model-package",
	Short: "Creates a model package that you can use to create SageMaker models or list on Amazon Web Services Marketplace, or a versioned model that is part of a model group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createModelPackageCmd).Standalone()

		sagemaker_createModelPackageCmd.Flags().String("additional-inference-specifications", "", "An array of additional Inference Specification objects.")
		sagemaker_createModelPackageCmd.Flags().String("certify-for-marketplace", "", "Whether to certify the model package for listing on Amazon Web Services Marketplace.")
		sagemaker_createModelPackageCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
		sagemaker_createModelPackageCmd.Flags().String("customer-metadata-properties", "", "The metadata properties associated with the model package versions.")
		sagemaker_createModelPackageCmd.Flags().String("domain", "", "The machine learning domain of your model package and its components.")
		sagemaker_createModelPackageCmd.Flags().String("drift-check-baselines", "", "Represents the drift check baselines that can be used when the model monitor is set using the model package.")
		sagemaker_createModelPackageCmd.Flags().String("inference-specification", "", "Specifies details about inference jobs that you can run with models based on this model package, including the following information:")
		sagemaker_createModelPackageCmd.Flags().String("metadata-properties", "", "")
		sagemaker_createModelPackageCmd.Flags().String("model-approval-status", "", "Whether the model is approved for deployment.")
		sagemaker_createModelPackageCmd.Flags().String("model-card", "", "The model card associated with the model package.")
		sagemaker_createModelPackageCmd.Flags().String("model-life-cycle", "", "A structure describing the current state of the model in its life cycle.")
		sagemaker_createModelPackageCmd.Flags().String("model-metrics", "", "A structure that contains model metrics reports.")
		sagemaker_createModelPackageCmd.Flags().String("model-package-description", "", "A description of the model package.")
		sagemaker_createModelPackageCmd.Flags().String("model-package-group-name", "", "The name or Amazon Resource Name (ARN) of the model package group that this model version belongs to.")
		sagemaker_createModelPackageCmd.Flags().String("model-package-name", "", "The name of the model package.")
		sagemaker_createModelPackageCmd.Flags().String("sample-payload-url", "", "The Amazon Simple Storage Service (Amazon S3) path where the sample payload is stored.")
		sagemaker_createModelPackageCmd.Flags().String("security-config", "", "The KMS Key ID (`KMSKeyId`) used for encryption of model package information.")
		sagemaker_createModelPackageCmd.Flags().String("skip-model-validation", "", "Indicates if you want to skip model validation.")
		sagemaker_createModelPackageCmd.Flags().String("source-algorithm-specification", "", "Details about the algorithm that was used to create the model package.")
		sagemaker_createModelPackageCmd.Flags().String("source-uri", "", "The URI of the source for the model package.")
		sagemaker_createModelPackageCmd.Flags().String("tags", "", "A list of key value pairs associated with the model.")
		sagemaker_createModelPackageCmd.Flags().String("task", "", "The machine learning task your model package accomplishes.")
		sagemaker_createModelPackageCmd.Flags().String("validation-specification", "", "Specifies configurations for one or more transform jobs that SageMaker runs to test the model package.")
	})
	sagemakerCmd.AddCommand(sagemaker_createModelPackageCmd)
}
