package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_deleteWhatsAppMessageMediaCmd = &cobra.Command{
	Use:   "delete-whats-app-message-media",
	Short: "Delete a media object from the WhatsApp service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_deleteWhatsAppMessageMediaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_deleteWhatsAppMessageMediaCmd).Standalone()

		socialmessaging_deleteWhatsAppMessageMediaCmd.Flags().String("media-id", "", "The unique identifier of the media file to delete.")
		socialmessaging_deleteWhatsAppMessageMediaCmd.Flags().String("origination-phone-number-id", "", "The unique identifier of the originating phone number associated with the media.")
		socialmessaging_deleteWhatsAppMessageMediaCmd.MarkFlagRequired("media-id")
		socialmessaging_deleteWhatsAppMessageMediaCmd.MarkFlagRequired("origination-phone-number-id")
	})
	socialmessagingCmd.AddCommand(socialmessaging_deleteWhatsAppMessageMediaCmd)
}
