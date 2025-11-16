package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd = &cobra.Command{
	Use:   "create-whats-app-message-template-from-library",
	Short: "Creates a new WhatsApp message template using a template from Meta's template library.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd).Standalone()

		socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account to associate with this template.")
		socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd.Flags().String("meta-library-template", "", "The template configuration from Meta's library, including customizations for buttons and body text.")
		socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd.MarkFlagRequired("id")
		socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd.MarkFlagRequired("meta-library-template")
	})
	socialmessagingCmd.AddCommand(socialmessaging_createWhatsAppMessageTemplateFromLibraryCmd)
}
