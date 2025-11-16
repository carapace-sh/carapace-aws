package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_registerTypeCmd = &cobra.Command{
	Use:   "register-type",
	Short: "Registers an extension with the CloudFormation service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_registerTypeCmd).Standalone()

	cloudformation_registerTypeCmd.Flags().String("client-request-token", "", "A unique identifier that acts as an idempotency key for this registration request.")
	cloudformation_registerTypeCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume when invoking the extension.")
	cloudformation_registerTypeCmd.Flags().String("logging-config", "", "Specifies logging configuration information for an extension.")
	cloudformation_registerTypeCmd.Flags().String("schema-handler-package", "", "A URL to the S3 bucket that contains the extension project package that contains the necessary files for the extension you want to register.")
	cloudformation_registerTypeCmd.Flags().String("type", "", "The kind of extension.")
	cloudformation_registerTypeCmd.Flags().String("type-name", "", "The name of the extension being registered.")
	cloudformation_registerTypeCmd.MarkFlagRequired("schema-handler-package")
	cloudformation_registerTypeCmd.MarkFlagRequired("type-name")
	cloudformationCmd.AddCommand(cloudformation_registerTypeCmd)
}
