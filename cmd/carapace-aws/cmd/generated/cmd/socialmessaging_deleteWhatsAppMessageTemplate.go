package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_deleteWhatsAppMessageTemplateCmd = &cobra.Command{
	Use:   "delete-whats-app-message-template",
	Short: "Deletes a WhatsApp message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_deleteWhatsAppMessageTemplateCmd).Standalone()

	socialmessaging_deleteWhatsAppMessageTemplateCmd.Flags().String("delete-all-languages", "", "If true, deletes all language versions of the template.")
	socialmessaging_deleteWhatsAppMessageTemplateCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account associated with this template.")
	socialmessaging_deleteWhatsAppMessageTemplateCmd.Flags().String("meta-template-id", "", "The numeric ID of the template assigned by Meta.")
	socialmessaging_deleteWhatsAppMessageTemplateCmd.Flags().String("template-name", "", "The name of the template to delete.")
	socialmessaging_deleteWhatsAppMessageTemplateCmd.MarkFlagRequired("id")
	socialmessaging_deleteWhatsAppMessageTemplateCmd.MarkFlagRequired("template-name")
	socialmessagingCmd.AddCommand(socialmessaging_deleteWhatsAppMessageTemplateCmd)
}
