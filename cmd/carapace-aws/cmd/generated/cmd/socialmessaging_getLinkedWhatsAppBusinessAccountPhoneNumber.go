package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_getLinkedWhatsAppBusinessAccountPhoneNumberCmd = &cobra.Command{
	Use:   "get-linked-whats-app-business-account-phone-number",
	Short: "Use your WhatsApp phone number id to get the WABA account id and phone number details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_getLinkedWhatsAppBusinessAccountPhoneNumberCmd).Standalone()

	socialmessaging_getLinkedWhatsAppBusinessAccountPhoneNumberCmd.Flags().String("id", "", "The unique identifier of the phone number.")
	socialmessaging_getLinkedWhatsAppBusinessAccountPhoneNumberCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_getLinkedWhatsAppBusinessAccountPhoneNumberCmd)
}
