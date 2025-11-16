package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendBounceCmd = &cobra.Command{
	Use:   "send-bounce",
	Short: "Generates and sends a bounce message to the sender of an email you received through Amazon SES.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendBounceCmd).Standalone()

	ses_sendBounceCmd.Flags().String("bounce-sender", "", "The address to use in the \"From\" header of the bounce message.")
	ses_sendBounceCmd.Flags().String("bounce-sender-arn", "", "This parameter is used only for sending authorization.")
	ses_sendBounceCmd.Flags().String("bounced-recipient-info-list", "", "A list of recipients of the bounced message, including the information required to create the Delivery Status Notifications (DSNs) for the recipients.")
	ses_sendBounceCmd.Flags().String("explanation", "", "Human-readable text for the bounce message to explain the failure.")
	ses_sendBounceCmd.Flags().String("message-dsn", "", "Message-related DSN fields.")
	ses_sendBounceCmd.Flags().String("original-message-id", "", "The message ID of the message to be bounced.")
	ses_sendBounceCmd.MarkFlagRequired("bounce-sender")
	ses_sendBounceCmd.MarkFlagRequired("bounced-recipient-info-list")
	ses_sendBounceCmd.MarkFlagRequired("original-message-id")
	sesCmd.AddCommand(ses_sendBounceCmd)
}
