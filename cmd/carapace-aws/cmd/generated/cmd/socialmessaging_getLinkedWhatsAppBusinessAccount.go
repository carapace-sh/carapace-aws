package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_getLinkedWhatsAppBusinessAccountCmd = &cobra.Command{
	Use:   "get-linked-whats-app-business-account",
	Short: "Get the details of your linked WhatsApp Business Account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_getLinkedWhatsAppBusinessAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_getLinkedWhatsAppBusinessAccountCmd).Standalone()

		socialmessaging_getLinkedWhatsAppBusinessAccountCmd.Flags().String("id", "", "The unique identifier, from Amazon Web Services, of the linked WhatsApp Business Account.")
		socialmessaging_getLinkedWhatsAppBusinessAccountCmd.MarkFlagRequired("id")
	})
	socialmessagingCmd.AddCommand(socialmessaging_getLinkedWhatsAppBusinessAccountCmd)
}
