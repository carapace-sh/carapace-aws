package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_listLinkedWhatsAppBusinessAccountsCmd = &cobra.Command{
	Use:   "list-linked-whats-app-business-accounts",
	Short: "List all WhatsApp Business Accounts linked to your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_listLinkedWhatsAppBusinessAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_listLinkedWhatsAppBusinessAccountsCmd).Standalone()

		socialmessaging_listLinkedWhatsAppBusinessAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		socialmessaging_listLinkedWhatsAppBusinessAccountsCmd.Flags().String("next-token", "", "The next token for pagination.")
	})
	socialmessagingCmd.AddCommand(socialmessaging_listLinkedWhatsAppBusinessAccountsCmd)
}
