package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_importModelCmd = &cobra.Command{
	Use:   "import-model",
	Short: "Creates a new custom model that replicates a source custom model that you import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_importModelCmd).Standalone()

	comprehend_importModelCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend permission to use Amazon Key Management Service (KMS) to encrypt or decrypt the custom model.")
	comprehend_importModelCmd.Flags().String("model-kms-key-id", "", "ID for the KMS key that Amazon Comprehend uses to encrypt trained custom models.")
	comprehend_importModelCmd.Flags().String("model-name", "", "The name to assign to the custom model that is created in Amazon Comprehend by this import.")
	comprehend_importModelCmd.Flags().String("source-model-arn", "", "The Amazon Resource Name (ARN) of the custom model to import.")
	comprehend_importModelCmd.Flags().String("tags", "", "Tags to associate with the custom model that is created by this import.")
	comprehend_importModelCmd.Flags().String("version-name", "", "The version name given to the custom model that is created by this import.")
	comprehend_importModelCmd.MarkFlagRequired("source-model-arn")
	comprehendCmd.AddCommand(comprehend_importModelCmd)
}
