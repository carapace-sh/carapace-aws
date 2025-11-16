package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_getWhatsAppMessageTemplateCmd = &cobra.Command{
	Use:   "get-whats-app-message-template",
	Short: "Retrieves a specific WhatsApp message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_getWhatsAppMessageTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_getWhatsAppMessageTemplateCmd).Standalone()

		socialmessaging_getWhatsAppMessageTemplateCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account associated with this template.")
		socialmessaging_getWhatsAppMessageTemplateCmd.Flags().String("meta-template-id", "", "The numeric ID of the template assigned by Meta.")
		socialmessaging_getWhatsAppMessageTemplateCmd.MarkFlagRequired("id")
		socialmessaging_getWhatsAppMessageTemplateCmd.MarkFlagRequired("meta-template-id")
	})
	socialmessagingCmd.AddCommand(socialmessaging_getWhatsAppMessageTemplateCmd)
}
