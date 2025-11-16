package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createModelCustomizationJobCmd = &cobra.Command{
	Use:   "create-model-customization-job",
	Short: "Creates a fine-tuning job to customize a base model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createModelCustomizationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createModelCustomizationJobCmd).Standalone()

		bedrock_createModelCustomizationJobCmd.Flags().String("base-model-identifier", "", "Name of the base model.")
		bedrock_createModelCustomizationJobCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrock_createModelCustomizationJobCmd.Flags().String("custom-model-kms-key-id", "", "The custom model is encrypted at rest using this key.")
		bedrock_createModelCustomizationJobCmd.Flags().String("custom-model-name", "", "A name for the resulting custom model.")
		bedrock_createModelCustomizationJobCmd.Flags().String("custom-model-tags", "", "Tags to attach to the resulting custom model.")
		bedrock_createModelCustomizationJobCmd.Flags().String("customization-config", "", "The customization configuration for the model customization job.")
		bedrock_createModelCustomizationJobCmd.Flags().String("customization-type", "", "The customization type.")
		bedrock_createModelCustomizationJobCmd.Flags().String("hyper-parameters", "", "Parameters related to tuning the model.")
		bedrock_createModelCustomizationJobCmd.Flags().String("job-name", "", "A name for the fine-tuning job.")
		bedrock_createModelCustomizationJobCmd.Flags().String("job-tags", "", "Tags to attach to the job.")
		bedrock_createModelCustomizationJobCmd.Flags().String("output-data-config", "", "S3 location for the output data.")
		bedrock_createModelCustomizationJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM service role that Amazon Bedrock can assume to perform tasks on your behalf.")
		bedrock_createModelCustomizationJobCmd.Flags().String("training-data-config", "", "Information about the training dataset.")
		bedrock_createModelCustomizationJobCmd.Flags().String("validation-data-config", "", "Information about the validation dataset.")
		bedrock_createModelCustomizationJobCmd.Flags().String("vpc-config", "", "The configuration of the Virtual Private Cloud (VPC) that contains the resources that you're using for this job.")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("base-model-identifier")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("custom-model-name")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("job-name")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("output-data-config")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("role-arn")
		bedrock_createModelCustomizationJobCmd.MarkFlagRequired("training-data-config")
	})
	bedrockCmd.AddCommand(bedrock_createModelCustomizationJobCmd)
}
