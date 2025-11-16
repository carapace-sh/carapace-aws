package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateModelPackageCmd = &cobra.Command{
	Use:   "update-model-package",
	Short: "Updates a versioned model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateModelPackageCmd).Standalone()

	sagemaker_updateModelPackageCmd.Flags().String("additional-inference-specifications-to-add", "", "An array of additional Inference Specification objects to be added to the existing array additional Inference Specification.")
	sagemaker_updateModelPackageCmd.Flags().String("approval-description", "", "A description for the approval status of the model.")
	sagemaker_updateModelPackageCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
	sagemaker_updateModelPackageCmd.Flags().String("customer-metadata-properties", "", "The metadata properties associated with the model package versions.")
	sagemaker_updateModelPackageCmd.Flags().String("customer-metadata-properties-to-remove", "", "The metadata properties associated with the model package versions to remove.")
	sagemaker_updateModelPackageCmd.Flags().String("inference-specification", "", "Specifies details about inference jobs that you can run with models based on this model package, including the following information:")
	sagemaker_updateModelPackageCmd.Flags().String("model-approval-status", "", "The approval status of the model.")
	sagemaker_updateModelPackageCmd.Flags().String("model-card", "", "The model card associated with the model package.")
	sagemaker_updateModelPackageCmd.Flags().String("model-life-cycle", "", "A structure describing the current state of the model in its life cycle.")
	sagemaker_updateModelPackageCmd.Flags().String("model-package-arn", "", "The Amazon Resource Name (ARN) of the model package.")
	sagemaker_updateModelPackageCmd.Flags().String("source-uri", "", "The URI of the source for the model package.")
	sagemaker_updateModelPackageCmd.MarkFlagRequired("model-package-arn")
	sagemakerCmd.AddCommand(sagemaker_updateModelPackageCmd)
}
