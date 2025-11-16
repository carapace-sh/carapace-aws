package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendEmailCmd = &cobra.Command{
	Use:   "send-email",
	Short: "Composes an email message and immediately queues it for sending.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_sendEmailCmd).Standalone()

		ses_sendEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when you send an email using `SendEmail`.")
		ses_sendEmailCmd.Flags().String("destination", "", "The destination for this email, composed of To:, CC:, and BCC: fields.")
		ses_sendEmailCmd.Flags().String("message", "", "The message to be sent.")
		ses_sendEmailCmd.Flags().String("reply-to-addresses", "", "The reply-to email address(es) for the message.")
		ses_sendEmailCmd.Flags().String("return-path", "", "The email address that bounces and complaints are forwarded to when feedback forwarding is enabled.")
		ses_sendEmailCmd.Flags().String("return-path-arn", "", "This parameter is used only for sending authorization.")
		ses_sendEmailCmd.Flags().String("source", "", "The email address that is sending the email.")
		ses_sendEmailCmd.Flags().String("source-arn", "", "This parameter is used only for sending authorization.")
		ses_sendEmailCmd.Flags().String("tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using `SendEmail`.")
		ses_sendEmailCmd.MarkFlagRequired("destination")
		ses_sendEmailCmd.MarkFlagRequired("message")
		ses_sendEmailCmd.MarkFlagRequired("source")
	})
	sesCmd.AddCommand(ses_sendEmailCmd)
}
