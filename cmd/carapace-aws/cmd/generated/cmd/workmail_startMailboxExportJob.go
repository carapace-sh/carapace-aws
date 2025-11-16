package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_startMailboxExportJobCmd = &cobra.Command{
	Use:   "start-mailbox-export-job",
	Short: "Starts a mailbox export job to export MIME-format email messages and calendar items from the specified mailbox to the specified Amazon Simple Storage Service (Amazon S3) bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_startMailboxExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_startMailboxExportJobCmd).Standalone()

		workmail_startMailboxExportJobCmd.Flags().String("client-token", "", "The idempotency token for the client request.")
		workmail_startMailboxExportJobCmd.Flags().String("description", "", "The mailbox export job description.")
		workmail_startMailboxExportJobCmd.Flags().String("entity-id", "", "The identifier of the user or resource associated with the mailbox.")
		workmail_startMailboxExportJobCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the symmetric AWS Key Management Service (AWS KMS) key that encrypts the exported mailbox content.")
		workmail_startMailboxExportJobCmd.Flags().String("organization-id", "", "The identifier associated with the organization.")
		workmail_startMailboxExportJobCmd.Flags().String("role-arn", "", "The ARN of the AWS Identity and Access Management (IAM) role that grants write permission to the S3 bucket.")
		workmail_startMailboxExportJobCmd.Flags().String("s3-bucket-name", "", "The name of the S3 bucket.")
		workmail_startMailboxExportJobCmd.Flags().String("s3-prefix", "", "The S3 bucket prefix.")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("client-token")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("entity-id")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("kms-key-arn")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("organization-id")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("role-arn")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("s3-bucket-name")
		workmail_startMailboxExportJobCmd.MarkFlagRequired("s3-prefix")
	})
	workmailCmd.AddCommand(workmail_startMailboxExportJobCmd)
}
