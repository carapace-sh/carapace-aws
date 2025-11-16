package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_sendWhatsAppMessageCmd = &cobra.Command{
	Use:   "send-whats-app-message",
	Short: "Send a WhatsApp message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_sendWhatsAppMessageCmd).Standalone()

	socialmessaging_sendWhatsAppMessageCmd.Flags().String("message", "", "The message to send through WhatsApp.")
	socialmessaging_sendWhatsAppMessageCmd.Flags().String("meta-api-version", "", "The API version for the request formatted as `v{VersionNumber}`.")
	socialmessaging_sendWhatsAppMessageCmd.Flags().String("origination-phone-number-id", "", "The ID of the phone number used to send the WhatsApp message.")
	socialmessaging_sendWhatsAppMessageCmd.MarkFlagRequired("message")
	socialmessaging_sendWhatsAppMessageCmd.MarkFlagRequired("meta-api-version")
	socialmessaging_sendWhatsAppMessageCmd.MarkFlagRequired("origination-phone-number-id")
	socialmessagingCmd.AddCommand(socialmessaging_sendWhatsAppMessageCmd)
}
