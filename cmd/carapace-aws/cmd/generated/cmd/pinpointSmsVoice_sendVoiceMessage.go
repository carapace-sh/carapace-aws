package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_sendVoiceMessageCmd = &cobra.Command{
	Use:   "send-voice-message",
	Short: "Create a new voice message and send it to a recipient's phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_sendVoiceMessageCmd).Standalone()

	pinpointSmsVoice_sendVoiceMessageCmd.Flags().String("caller-id", "", "The phone number that appears on recipients' devices when they receive the message.")
	pinpointSmsVoice_sendVoiceMessageCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to use to send the message.")
	pinpointSmsVoice_sendVoiceMessageCmd.Flags().String("content", "", "")
	pinpointSmsVoice_sendVoiceMessageCmd.Flags().String("destination-phone-number", "", "The phone number that you want to send the voice message to.")
	pinpointSmsVoice_sendVoiceMessageCmd.Flags().String("origination-phone-number", "", "The phone number that Amazon Pinpoint should use to send the voice message.")
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_sendVoiceMessageCmd)
}
