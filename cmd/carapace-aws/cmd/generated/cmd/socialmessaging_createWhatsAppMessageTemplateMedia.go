package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_createWhatsAppMessageTemplateMediaCmd = &cobra.Command{
	Use:   "create-whats-app-message-template-media",
	Short: "Uploads media for use in a WhatsApp message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_createWhatsAppMessageTemplateMediaCmd).Standalone()

	socialmessaging_createWhatsAppMessageTemplateMediaCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account associated with this media upload.")
	socialmessaging_createWhatsAppMessageTemplateMediaCmd.Flags().String("source-s3-file", "", "")
	socialmessaging_createWhatsAppMessageTemplateMediaCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_createWhatsAppMessageTemplateMediaCmd)
}
