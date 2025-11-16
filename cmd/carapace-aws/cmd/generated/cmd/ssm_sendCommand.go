package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_sendCommandCmd = &cobra.Command{
	Use:   "send-command",
	Short: "Runs commands on one or more managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_sendCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_sendCommandCmd).Standalone()

		ssm_sendCommandCmd.Flags().String("alarm-configuration", "", "The CloudWatch alarm you want to apply to your command.")
		ssm_sendCommandCmd.Flags().String("cloud-watch-output-config", "", "Enables Amazon Web Services Systems Manager to send Run Command output to Amazon CloudWatch Logs.")
		ssm_sendCommandCmd.Flags().String("comment", "", "User-specified information about the command, such as a brief description of what the command should do.")
		ssm_sendCommandCmd.Flags().String("document-hash", "", "The Sha256 or Sha1 hash created by the system when the document was created.")
		ssm_sendCommandCmd.Flags().String("document-hash-type", "", "Sha256 or Sha1.")
		ssm_sendCommandCmd.Flags().String("document-name", "", "The name of the Amazon Web Services Systems Manager document (SSM document) to run.")
		ssm_sendCommandCmd.Flags().String("document-version", "", "The SSM document version to use in the request.")
		ssm_sendCommandCmd.Flags().String("instance-ids", "", "The IDs of the managed nodes where the command should run.")
		ssm_sendCommandCmd.Flags().String("max-concurrency", "", "(Optional) The maximum number of managed nodes that are allowed to run the command at the same time.")
		ssm_sendCommandCmd.Flags().String("max-errors", "", "The maximum number of errors allowed without the command failing.")
		ssm_sendCommandCmd.Flags().String("notification-config", "", "Configurations for sending notifications.")
		ssm_sendCommandCmd.Flags().String("output-s3-bucket-name", "", "The name of the S3 bucket where command execution responses should be stored.")
		ssm_sendCommandCmd.Flags().String("output-s3-key-prefix", "", "The directory structure within the S3 bucket where the responses should be stored.")
		ssm_sendCommandCmd.Flags().String("output-s3-region", "", "(Deprecated) You can no longer specify this parameter.")
		ssm_sendCommandCmd.Flags().String("parameters", "", "The required and optional parameters specified in the document being run.")
		ssm_sendCommandCmd.Flags().String("service-role-arn", "", "The ARN of the Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for Run Command commands.")
		ssm_sendCommandCmd.Flags().String("targets", "", "An array of search criteria that targets managed nodes using a `Key,Value` combination that you specify.")
		ssm_sendCommandCmd.Flags().String("timeout-seconds", "", "If this time is reached and the command hasn't already started running, it won't run.")
		ssm_sendCommandCmd.MarkFlagRequired("document-name")
	})
	ssmCmd.AddCommand(ssm_sendCommandCmd)
}
