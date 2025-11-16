package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_copyProjectVersionCmd = &cobra.Command{
	Use:   "copy-project-version",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_copyProjectVersionCmd).Standalone()

	rekognition_copyProjectVersionCmd.Flags().String("destination-project-arn", "", "The ARN of the project in the trusted AWS account that you want to copy the model version to.")
	rekognition_copyProjectVersionCmd.Flags().String("kms-key-id", "", "The identifier for your AWS Key Management Service key (AWS KMS key).")
	rekognition_copyProjectVersionCmd.Flags().String("output-config", "", "The S3 bucket and folder location where the training output for the source model version is placed.")
	rekognition_copyProjectVersionCmd.Flags().String("source-project-arn", "", "The ARN of the source project in the trusting AWS account.")
	rekognition_copyProjectVersionCmd.Flags().String("source-project-version-arn", "", "The ARN of the model version in the source project that you want to copy to a destination project.")
	rekognition_copyProjectVersionCmd.Flags().String("tags", "", "The key-value tags to assign to the model version.")
	rekognition_copyProjectVersionCmd.Flags().String("version-name", "", "A name for the version of the model that's copied to the destination project.")
	rekognition_copyProjectVersionCmd.MarkFlagRequired("destination-project-arn")
	rekognition_copyProjectVersionCmd.MarkFlagRequired("output-config")
	rekognition_copyProjectVersionCmd.MarkFlagRequired("source-project-arn")
	rekognition_copyProjectVersionCmd.MarkFlagRequired("source-project-version-arn")
	rekognition_copyProjectVersionCmd.MarkFlagRequired("version-name")
	rekognitionCmd.AddCommand(rekognition_copyProjectVersionCmd)
}
