package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendBulkTemplatedEmailCmd = &cobra.Command{
	Use:   "send-bulk-templated-email",
	Short: "Composes an email message to multiple destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendBulkTemplatedEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_sendBulkTemplatedEmailCmd).Standalone()

		ses_sendBulkTemplatedEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when you send an email using `SendBulkTemplatedEmail`.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("default-tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send to a destination using `SendBulkTemplatedEmail`.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("default-template-data", "", "A list of replacement values to apply to the template when replacement data is not specified in a Destination object.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("destinations", "", "One or more `Destination` objects.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("reply-to-addresses", "", "The reply-to email address(es) for the message.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("return-path", "", "The email address that bounces and complaints are forwarded to when feedback forwarding is enabled.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("return-path-arn", "", "This parameter is used only for sending authorization.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("source", "", "The email address that is sending the email.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("source-arn", "", "This parameter is used only for sending authorization.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("template", "", "The template to use when sending this email.")
		ses_sendBulkTemplatedEmailCmd.Flags().String("template-arn", "", "The ARN of the template to use when sending this email.")
		ses_sendBulkTemplatedEmailCmd.MarkFlagRequired("default-template-data")
		ses_sendBulkTemplatedEmailCmd.MarkFlagRequired("destinations")
		ses_sendBulkTemplatedEmailCmd.MarkFlagRequired("source")
		ses_sendBulkTemplatedEmailCmd.MarkFlagRequired("template")
	})
	sesCmd.AddCommand(ses_sendBulkTemplatedEmailCmd)
}
