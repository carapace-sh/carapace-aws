package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createProjectVersionCmd = &cobra.Command{
	Use:   "create-project-version",
	Short: "Creates a new version of Amazon Rekognition project (like a Custom Labels model or a custom adapter) and begins training.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createProjectVersionCmd).Standalone()

	rekognition_createProjectVersionCmd.Flags().String("feature-config", "", "Feature-specific configuration of the training job.")
	rekognition_createProjectVersionCmd.Flags().String("kms-key-id", "", "The identifier for your AWS Key Management Service key (AWS KMS key).")
	rekognition_createProjectVersionCmd.Flags().String("output-config", "", "The Amazon S3 bucket location to store the results of training.")
	rekognition_createProjectVersionCmd.Flags().String("project-arn", "", "The ARN of the Amazon Rekognition project that will manage the project version you want to train.")
	rekognition_createProjectVersionCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the project version.")
	rekognition_createProjectVersionCmd.Flags().String("testing-data", "", "Specifies an external manifest that the service uses to test the project version.")
	rekognition_createProjectVersionCmd.Flags().String("training-data", "", "Specifies an external manifest that the services uses to train the project version.")
	rekognition_createProjectVersionCmd.Flags().String("version-description", "", "A description applied to the project version being created.")
	rekognition_createProjectVersionCmd.Flags().String("version-name", "", "A name for the version of the project version.")
	rekognition_createProjectVersionCmd.MarkFlagRequired("output-config")
	rekognition_createProjectVersionCmd.MarkFlagRequired("project-arn")
	rekognition_createProjectVersionCmd.MarkFlagRequired("version-name")
	rekognitionCmd.AddCommand(rekognition_createProjectVersionCmd)
}
