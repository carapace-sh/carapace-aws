package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createModelCopyJobCmd = &cobra.Command{
	Use:   "create-model-copy-job",
	Short: "Copies a model to another region so that it can be used there.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createModelCopyJobCmd).Standalone()

	bedrock_createModelCopyJobCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrock_createModelCopyJobCmd.Flags().String("model-kms-key-id", "", "The ARN of the KMS key that you use to encrypt the model copy.")
	bedrock_createModelCopyJobCmd.Flags().String("source-model-arn", "", "The Amazon Resource Name (ARN) of the model to be copied.")
	bedrock_createModelCopyJobCmd.Flags().String("target-model-name", "", "A name for the copied model.")
	bedrock_createModelCopyJobCmd.Flags().String("target-model-tags", "", "Tags to associate with the target model.")
	bedrock_createModelCopyJobCmd.MarkFlagRequired("source-model-arn")
	bedrock_createModelCopyJobCmd.MarkFlagRequired("target-model-name")
	bedrockCmd.AddCommand(bedrock_createModelCopyJobCmd)
}
