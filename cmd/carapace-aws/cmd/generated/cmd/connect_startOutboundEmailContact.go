package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startOutboundEmailContactCmd = &cobra.Command{
	Use:   "start-outbound-email-contact",
	Short: "Initiates a flow to send an agent reply or outbound email contact (created from the CreateContact API) to a customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startOutboundEmailContactCmd).Standalone()

	connect_startOutboundEmailContactCmd.Flags().String("additional-recipients", "", "The additional recipients address of email in CC.")
	connect_startOutboundEmailContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startOutboundEmailContactCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
	connect_startOutboundEmailContactCmd.Flags().String("destination-email-address", "", "The email address of the customer.")
	connect_startOutboundEmailContactCmd.Flags().String("email-message", "", "The email message body to be sent to the newly created email.")
	connect_startOutboundEmailContactCmd.Flags().String("from-email-address", "", "The email address associated with the Amazon Connect instance.")
	connect_startOutboundEmailContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startOutboundEmailContactCmd.MarkFlagRequired("contact-id")
	connect_startOutboundEmailContactCmd.MarkFlagRequired("destination-email-address")
	connect_startOutboundEmailContactCmd.MarkFlagRequired("email-message")
	connect_startOutboundEmailContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_startOutboundEmailContactCmd)
}
