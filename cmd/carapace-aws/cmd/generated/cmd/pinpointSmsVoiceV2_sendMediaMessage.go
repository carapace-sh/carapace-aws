package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_sendMediaMessageCmd = &cobra.Command{
	Use:   "send-media-message",
	Short: "Creates a new multimedia message (MMS) and sends it to a recipient's phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_sendMediaMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_sendMediaMessageCmd).Standalone()

		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("context", "", "You can specify custom data in this field.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("destination-phone-number", "", "The destination phone number in E.164 format.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("dry-run", "", "When set to true, the message is checked and validated, but isn't sent to the end recipient.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("max-price", "", "The maximum amount that you want to spend, in US dollars, per each MMS message.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("media-urls", "", "An array of URLs to each media file to send.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("message-body", "", "The text body of the message.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().Bool("message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().Bool("no-message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("origination-identity", "", "The origination identity of the message.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("protect-configuration-id", "", "The unique identifier of the protect configuration to use.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flags().String("time-to-live", "", "How long the media message is valid for.")
		pinpointSmsVoiceV2_sendMediaMessageCmd.MarkFlagRequired("destination-phone-number")
		pinpointSmsVoiceV2_sendMediaMessageCmd.Flag("no-message-feedback-enabled").Hidden = true
		pinpointSmsVoiceV2_sendMediaMessageCmd.MarkFlagRequired("origination-identity")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_sendMediaMessageCmd)
}
