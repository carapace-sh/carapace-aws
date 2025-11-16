package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_createDocumentClassifierCmd = &cobra.Command{
	Use:   "create-document-classifier",
	Short: "Creates a new document classifier that you can use to categorize documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_createDocumentClassifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_createDocumentClassifierCmd).Standalone()

		comprehend_createDocumentClassifierCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_createDocumentClassifierCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_createDocumentClassifierCmd.Flags().String("document-classifier-name", "", "The name of the document classifier.")
		comprehend_createDocumentClassifierCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
		comprehend_createDocumentClassifierCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_createDocumentClassifierCmd.Flags().String("mode", "", "Indicates the mode in which the classifier will be trained.")
		comprehend_createDocumentClassifierCmd.Flags().String("model-kms-key-id", "", "ID for the KMS key that Amazon Comprehend uses to encrypt trained custom models.")
		comprehend_createDocumentClassifierCmd.Flags().String("model-policy", "", "The resource-based policy to attach to your custom document classifier model.")
		comprehend_createDocumentClassifierCmd.Flags().String("output-data-config", "", "Specifies the location for the output files from a custom classifier job.")
		comprehend_createDocumentClassifierCmd.Flags().String("tags", "", "Tags to associate with the document classifier.")
		comprehend_createDocumentClassifierCmd.Flags().String("version-name", "", "The version name given to the newly created classifier.")
		comprehend_createDocumentClassifierCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
		comprehend_createDocumentClassifierCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your custom classifier.")
		comprehend_createDocumentClassifierCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_createDocumentClassifierCmd.MarkFlagRequired("document-classifier-name")
		comprehend_createDocumentClassifierCmd.MarkFlagRequired("input-data-config")
		comprehend_createDocumentClassifierCmd.MarkFlagRequired("language-code")
	})
	comprehendCmd.AddCommand(comprehend_createDocumentClassifierCmd)
}
