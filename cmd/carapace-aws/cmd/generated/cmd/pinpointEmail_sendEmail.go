package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_sendEmailCmd = &cobra.Command{
	Use:   "send-email",
	Short: "Sends an email message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_sendEmailCmd).Standalone()

	pinpointEmail_sendEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to use when sending the email.")
	pinpointEmail_sendEmailCmd.Flags().String("content", "", "An object that contains the body of the message.")
	pinpointEmail_sendEmailCmd.Flags().String("destination", "", "An object that contains the recipients of the email message.")
	pinpointEmail_sendEmailCmd.Flags().String("email-tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using the `SendEmail` operation.")
	pinpointEmail_sendEmailCmd.Flags().String("feedback-forwarding-email-address", "", "The address that Amazon Pinpoint should send bounce and complaint notifications to.")
	pinpointEmail_sendEmailCmd.Flags().String("from-email-address", "", "The email address that you want to use as the \"From\" address for the email.")
	pinpointEmail_sendEmailCmd.Flags().String("reply-to-addresses", "", "The \"Reply-to\" email addresses for the message.")
	pinpointEmail_sendEmailCmd.MarkFlagRequired("content")
	pinpointEmail_sendEmailCmd.MarkFlagRequired("destination")
	pinpointEmailCmd.AddCommand(pinpointEmail_sendEmailCmd)
}
