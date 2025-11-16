package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startDocumentClassificationJobCmd = &cobra.Command{
	Use:   "start-document-classification-job",
	Short: "Starts an asynchronous document classification job using a custom classification model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startDocumentClassificationJobCmd).Standalone()

	comprehend_startDocumentClassificationJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("document-classifier-arn", "", "The Amazon Resource Name (ARN) of the document classifier to use to process the job.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel associated with the model to use.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("job-name", "", "The identifier of the job.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("tags", "", "Tags to associate with the document classification job.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
	comprehend_startDocumentClassificationJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your document classification job.")
	comprehend_startDocumentClassificationJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_startDocumentClassificationJobCmd.MarkFlagRequired("input-data-config")
	comprehend_startDocumentClassificationJobCmd.MarkFlagRequired("output-data-config")
	comprehendCmd.AddCommand(comprehend_startDocumentClassificationJobCmd)
}
