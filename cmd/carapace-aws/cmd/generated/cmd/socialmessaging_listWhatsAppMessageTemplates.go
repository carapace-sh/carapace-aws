package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_listWhatsAppMessageTemplatesCmd = &cobra.Command{
	Use:   "list-whats-app-message-templates",
	Short: "Lists WhatsApp message templates for a specific WhatsApp Business Account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_listWhatsAppMessageTemplatesCmd).Standalone()

	socialmessaging_listWhatsAppMessageTemplatesCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account to list templates for.")
	socialmessaging_listWhatsAppMessageTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page (1-100).")
	socialmessaging_listWhatsAppMessageTemplatesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	socialmessaging_listWhatsAppMessageTemplatesCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_listWhatsAppMessageTemplatesCmd)
}
