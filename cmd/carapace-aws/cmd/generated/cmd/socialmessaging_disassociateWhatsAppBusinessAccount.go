package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_disassociateWhatsAppBusinessAccountCmd = &cobra.Command{
	Use:   "disassociate-whats-app-business-account",
	Short: "Disassociate a WhatsApp Business Account (WABA) from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_disassociateWhatsAppBusinessAccountCmd).Standalone()

	socialmessaging_disassociateWhatsAppBusinessAccountCmd.Flags().String("id", "", "The unique identifier of your WhatsApp Business Account.")
	socialmessaging_disassociateWhatsAppBusinessAccountCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_disassociateWhatsAppBusinessAccountCmd)
}
