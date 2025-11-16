package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_sendEmailCmd = &cobra.Command{
	Use:   "send-email",
	Short: "Sends an email message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_sendEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_sendEmailCmd).Standalone()

		sesv2_sendEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when sending the email.")
		sesv2_sendEmailCmd.Flags().String("content", "", "An object that contains the body of the message.")
		sesv2_sendEmailCmd.Flags().String("destination", "", "An object that contains the recipients of the email message.")
		sesv2_sendEmailCmd.Flags().String("email-tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using the `SendEmail` operation.")
		sesv2_sendEmailCmd.Flags().String("endpoint-id", "", "The ID of the multi-region endpoint (global-endpoint).")
		sesv2_sendEmailCmd.Flags().String("feedback-forwarding-email-address", "", "The address that you want bounce and complaint notifications to be sent to.")
		sesv2_sendEmailCmd.Flags().String("feedback-forwarding-email-address-identity-arn", "", "This parameter is used only for sending authorization.")
		sesv2_sendEmailCmd.Flags().String("from-email-address", "", "The email address to use as the \"From\" address for the email.")
		sesv2_sendEmailCmd.Flags().String("from-email-address-identity-arn", "", "This parameter is used only for sending authorization.")
		sesv2_sendEmailCmd.Flags().String("list-management-options", "", "An object used to specify a list or topic to which an email belongs, which will be used when a contact chooses to unsubscribe.")
		sesv2_sendEmailCmd.Flags().String("reply-to-addresses", "", "The \"Reply-to\" email addresses for the message.")
		sesv2_sendEmailCmd.Flags().String("tenant-name", "", "The name of the tenant through which this email will be sent.")
		sesv2_sendEmailCmd.MarkFlagRequired("content")
	})
	sesv2Cmd.AddCommand(sesv2_sendEmailCmd)
}
