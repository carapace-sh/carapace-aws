package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendTemplatedEmailCmd = &cobra.Command{
	Use:   "send-templated-email",
	Short: "Composes an email message using an email template and immediately queues it for sending.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendTemplatedEmailCmd).Standalone()

	ses_sendTemplatedEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when you send an email using `SendTemplatedEmail`.")
	ses_sendTemplatedEmailCmd.Flags().String("destination", "", "The destination for this email, composed of To:, CC:, and BCC: fields.")
	ses_sendTemplatedEmailCmd.Flags().String("reply-to-addresses", "", "The reply-to email address(es) for the message.")
	ses_sendTemplatedEmailCmd.Flags().String("return-path", "", "The email address that bounces and complaints are forwarded to when feedback forwarding is enabled.")
	ses_sendTemplatedEmailCmd.Flags().String("return-path-arn", "", "This parameter is used only for sending authorization.")
	ses_sendTemplatedEmailCmd.Flags().String("source", "", "The email address that is sending the email.")
	ses_sendTemplatedEmailCmd.Flags().String("source-arn", "", "This parameter is used only for sending authorization.")
	ses_sendTemplatedEmailCmd.Flags().String("tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using `SendTemplatedEmail`.")
	ses_sendTemplatedEmailCmd.Flags().String("template", "", "The template to use when sending this email.")
	ses_sendTemplatedEmailCmd.Flags().String("template-arn", "", "The ARN of the template to use when sending this email.")
	ses_sendTemplatedEmailCmd.Flags().String("template-data", "", "A list of replacement values to apply to the template.")
	ses_sendTemplatedEmailCmd.MarkFlagRequired("destination")
	ses_sendTemplatedEmailCmd.MarkFlagRequired("source")
	ses_sendTemplatedEmailCmd.MarkFlagRequired("template")
	ses_sendTemplatedEmailCmd.MarkFlagRequired("template-data")
	sesCmd.AddCommand(ses_sendTemplatedEmailCmd)
}
