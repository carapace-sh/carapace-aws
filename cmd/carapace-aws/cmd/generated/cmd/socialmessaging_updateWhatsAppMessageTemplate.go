package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_updateWhatsAppMessageTemplateCmd = &cobra.Command{
	Use:   "update-whats-app-message-template",
	Short: "Updates an existing WhatsApp message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_updateWhatsAppMessageTemplateCmd).Standalone()

	socialmessaging_updateWhatsAppMessageTemplateCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account associated with this template.")
	socialmessaging_updateWhatsAppMessageTemplateCmd.Flags().String("meta-template-id", "", "The numeric ID of the template assigned by Meta.")
	socialmessaging_updateWhatsAppMessageTemplateCmd.Flags().String("template-category", "", "The new category for the template (for example, UTILITY or MARKETING).")
	socialmessaging_updateWhatsAppMessageTemplateCmd.Flags().String("template-components", "", "The updated components of the template as a JSON blob (maximum 3000 characters).")
	socialmessaging_updateWhatsAppMessageTemplateCmd.MarkFlagRequired("id")
	socialmessaging_updateWhatsAppMessageTemplateCmd.MarkFlagRequired("meta-template-id")
	socialmessagingCmd.AddCommand(socialmessaging_updateWhatsAppMessageTemplateCmd)
}
