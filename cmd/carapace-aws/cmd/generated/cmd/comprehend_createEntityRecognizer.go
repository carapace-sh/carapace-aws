package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_createEntityRecognizerCmd = &cobra.Command{
	Use:   "create-entity-recognizer",
	Short: "Creates an entity recognizer using submitted files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_createEntityRecognizerCmd).Standalone()

	comprehend_createEntityRecognizerCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehend_createEntityRecognizerCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
	comprehend_createEntityRecognizerCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data.")
	comprehend_createEntityRecognizerCmd.Flags().String("language-code", "", "You can specify any of the following languages: English (\"en\"), Spanish (\"es\"), French (\"fr\"), Italian (\"it\"), German (\"de\"), or Portuguese (\"pt\").")
	comprehend_createEntityRecognizerCmd.Flags().String("model-kms-key-id", "", "ID for the KMS key that Amazon Comprehend uses to encrypt trained custom models.")
	comprehend_createEntityRecognizerCmd.Flags().String("model-policy", "", "The JSON resource-based policy to attach to your custom entity recognizer model.")
	comprehend_createEntityRecognizerCmd.Flags().String("recognizer-name", "", "The name given to the newly created recognizer.")
	comprehend_createEntityRecognizerCmd.Flags().String("tags", "", "Tags to associate with the entity recognizer.")
	comprehend_createEntityRecognizerCmd.Flags().String("version-name", "", "The version name given to the newly created recognizer.")
	comprehend_createEntityRecognizerCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
	comprehend_createEntityRecognizerCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your custom entity recognizer.")
	comprehend_createEntityRecognizerCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_createEntityRecognizerCmd.MarkFlagRequired("input-data-config")
	comprehend_createEntityRecognizerCmd.MarkFlagRequired("language-code")
	comprehend_createEntityRecognizerCmd.MarkFlagRequired("recognizer-name")
	comprehendCmd.AddCommand(comprehend_createEntityRecognizerCmd)
}
