package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_getWhatsAppMessageMediaCmd = &cobra.Command{
	Use:   "get-whats-app-message-media",
	Short: "Get a media file from the WhatsApp service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_getWhatsAppMessageMediaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_getWhatsAppMessageMediaCmd).Standalone()

		socialmessaging_getWhatsAppMessageMediaCmd.Flags().String("destination-s3-file", "", "The `bucketName` and `key` of the S3 media file.")
		socialmessaging_getWhatsAppMessageMediaCmd.Flags().String("destination-s3-presigned-url", "", "The presign url of the media file.")
		socialmessaging_getWhatsAppMessageMediaCmd.Flags().String("media-id", "", "The unique identifier for the media file.")
		socialmessaging_getWhatsAppMessageMediaCmd.Flags().Bool("metadata-only", false, "Set to `True` to get only the metadata for the file.")
		socialmessaging_getWhatsAppMessageMediaCmd.Flags().Bool("no-metadata-only", false, "Set to `True` to get only the metadata for the file.")
		socialmessaging_getWhatsAppMessageMediaCmd.Flags().String("origination-phone-number-id", "", "The unique identifier of the originating phone number for the WhatsApp message media.")
		socialmessaging_getWhatsAppMessageMediaCmd.MarkFlagRequired("media-id")
		socialmessaging_getWhatsAppMessageMediaCmd.Flag("no-metadata-only").Hidden = true
		socialmessaging_getWhatsAppMessageMediaCmd.MarkFlagRequired("origination-phone-number-id")
	})
	socialmessagingCmd.AddCommand(socialmessaging_getWhatsAppMessageMediaCmd)
}
