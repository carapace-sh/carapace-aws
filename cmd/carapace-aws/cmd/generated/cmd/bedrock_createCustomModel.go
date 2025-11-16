package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createCustomModelCmd = &cobra.Command{
	Use:   "create-custom-model",
	Short: "Creates a new custom model in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createCustomModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createCustomModelCmd).Standalone()

		bedrock_createCustomModelCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrock_createCustomModelCmd.Flags().String("model-kms-key-arn", "", "The Amazon Resource Name (ARN) of the customer managed KMS key to encrypt the custom model.")
		bedrock_createCustomModelCmd.Flags().String("model-name", "", "A unique name for the custom model.")
		bedrock_createCustomModelCmd.Flags().String("model-source-config", "", "The data source for the model.")
		bedrock_createCustomModelCmd.Flags().String("model-tags", "", "A list of key-value pairs to associate with the custom model resource.")
		bedrock_createCustomModelCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM service role that Amazon Bedrock assumes to perform tasks on your behalf.")
		bedrock_createCustomModelCmd.MarkFlagRequired("model-name")
		bedrock_createCustomModelCmd.MarkFlagRequired("model-source-config")
	})
	bedrockCmd.AddCommand(bedrock_createCustomModelCmd)
}
