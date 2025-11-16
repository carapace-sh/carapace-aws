package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_postWhatsAppMessageMediaCmd = &cobra.Command{
	Use:   "post-whats-app-message-media",
	Short: "Upload a media file to the WhatsApp service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_postWhatsAppMessageMediaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_postWhatsAppMessageMediaCmd).Standalone()

		socialmessaging_postWhatsAppMessageMediaCmd.Flags().String("origination-phone-number-id", "", "The ID of the phone number to associate with the WhatsApp media file.")
		socialmessaging_postWhatsAppMessageMediaCmd.Flags().String("source-s3-file", "", "The source S3 url for the media file.")
		socialmessaging_postWhatsAppMessageMediaCmd.Flags().String("source-s3-presigned-url", "", "The source presign url of the media file.")
		socialmessaging_postWhatsAppMessageMediaCmd.MarkFlagRequired("origination-phone-number-id")
	})
	socialmessagingCmd.AddCommand(socialmessaging_postWhatsAppMessageMediaCmd)
}
