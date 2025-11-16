package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_sendBulkEmailCmd = &cobra.Command{
	Use:   "send-bulk-email",
	Short: "Composes an email message to multiple destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_sendBulkEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_sendBulkEmailCmd).Standalone()

		sesv2_sendBulkEmailCmd.Flags().String("bulk-email-entries", "", "The list of bulk email entry objects.")
		sesv2_sendBulkEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when sending the email.")
		sesv2_sendBulkEmailCmd.Flags().String("default-content", "", "An object that contains the body of the message.")
		sesv2_sendBulkEmailCmd.Flags().String("default-email-tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using the `SendEmail` operation.")
		sesv2_sendBulkEmailCmd.Flags().String("endpoint-id", "", "The ID of the multi-region endpoint (global-endpoint).")
		sesv2_sendBulkEmailCmd.Flags().String("feedback-forwarding-email-address", "", "The address that you want bounce and complaint notifications to be sent to.")
		sesv2_sendBulkEmailCmd.Flags().String("feedback-forwarding-email-address-identity-arn", "", "This parameter is used only for sending authorization.")
		sesv2_sendBulkEmailCmd.Flags().String("from-email-address", "", "The email address to use as the \"From\" address for the email.")
		sesv2_sendBulkEmailCmd.Flags().String("from-email-address-identity-arn", "", "This parameter is used only for sending authorization.")
		sesv2_sendBulkEmailCmd.Flags().String("reply-to-addresses", "", "The \"Reply-to\" email addresses for the message.")
		sesv2_sendBulkEmailCmd.Flags().String("tenant-name", "", "The name of the tenant through which this bulk email will be sent.")
		sesv2_sendBulkEmailCmd.MarkFlagRequired("bulk-email-entries")
		sesv2_sendBulkEmailCmd.MarkFlagRequired("default-content")
	})
	sesv2Cmd.AddCommand(sesv2_sendBulkEmailCmd)
}
