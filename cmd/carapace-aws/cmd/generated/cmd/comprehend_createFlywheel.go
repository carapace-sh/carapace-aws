package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_createFlywheelCmd = &cobra.Command{
	Use:   "create-flywheel",
	Short: "A flywheel is an Amazon Web Services resource that orchestrates the ongoing training of a model for custom classification or custom entity recognition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_createFlywheelCmd).Standalone()

	comprehend_createFlywheelCmd.Flags().String("active-model-arn", "", "To associate an existing model with the flywheel, specify the Amazon Resource Number (ARN) of the model version.")
	comprehend_createFlywheelCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehend_createFlywheelCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend the permissions required to access the flywheel data in the data lake.")
	comprehend_createFlywheelCmd.Flags().String("data-lake-s3-uri", "", "Enter the S3 location for the data lake.")
	comprehend_createFlywheelCmd.Flags().String("data-security-config", "", "Data security configurations.")
	comprehend_createFlywheelCmd.Flags().String("flywheel-name", "", "Name for the flywheel.")
	comprehend_createFlywheelCmd.Flags().String("model-type", "", "The model type.")
	comprehend_createFlywheelCmd.Flags().String("tags", "", "The tags to associate with this flywheel.")
	comprehend_createFlywheelCmd.Flags().String("task-config", "", "Configuration about the model associated with the flywheel.")
	comprehend_createFlywheelCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_createFlywheelCmd.MarkFlagRequired("data-lake-s3-uri")
	comprehend_createFlywheelCmd.MarkFlagRequired("flywheel-name")
	comprehendCmd.AddCommand(comprehend_createFlywheelCmd)
}
