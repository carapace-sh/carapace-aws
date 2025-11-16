package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_sendOutboundEmailCmd = &cobra.Command{
	Use:   "send-outbound-email",
	Short: "Send outbound email for outbound campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_sendOutboundEmailCmd).Standalone()

	connect_sendOutboundEmailCmd.Flags().String("additional-recipients", "", "The additional recipients address of the email in CC.")
	connect_sendOutboundEmailCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_sendOutboundEmailCmd.Flags().String("destination-email-address", "", "The email address to send the email to.")
	connect_sendOutboundEmailCmd.Flags().String("email-message", "", "The email message body to be sent to the newly created email.")
	connect_sendOutboundEmailCmd.Flags().String("from-email-address", "", "The email address to be used for sending email.")
	connect_sendOutboundEmailCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_sendOutboundEmailCmd.Flags().String("source-campaign", "", "A Campaign object need for Campaign traffic type.")
	connect_sendOutboundEmailCmd.Flags().String("traffic-type", "", "Denotes the class of traffic.")
	connect_sendOutboundEmailCmd.MarkFlagRequired("destination-email-address")
	connect_sendOutboundEmailCmd.MarkFlagRequired("email-message")
	connect_sendOutboundEmailCmd.MarkFlagRequired("from-email-address")
	connect_sendOutboundEmailCmd.MarkFlagRequired("instance-id")
	connect_sendOutboundEmailCmd.MarkFlagRequired("traffic-type")
	connectCmd.AddCommand(connect_sendOutboundEmailCmd)
}
