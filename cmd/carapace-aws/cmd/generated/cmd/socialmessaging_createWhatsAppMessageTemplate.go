package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_createWhatsAppMessageTemplateCmd = &cobra.Command{
	Use:   "create-whats-app-message-template",
	Short: "Creates a new WhatsApp message template from a custom definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_createWhatsAppMessageTemplateCmd).Standalone()

	socialmessaging_createWhatsAppMessageTemplateCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account to associate with this template.")
	socialmessaging_createWhatsAppMessageTemplateCmd.Flags().String("template-definition", "", "The complete template definition as a JSON blob.")
	socialmessaging_createWhatsAppMessageTemplateCmd.MarkFlagRequired("id")
	socialmessaging_createWhatsAppMessageTemplateCmd.MarkFlagRequired("template-definition")
	socialmessagingCmd.AddCommand(socialmessaging_createWhatsAppMessageTemplateCmd)
}
