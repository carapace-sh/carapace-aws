package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd = &cobra.Command{
	Use:   "send-destination-number-verification-code",
	Short: "Before you can send test messages to a verified destination phone number you need to opt-in the verified destination phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd).Standalone()

	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("context", "", "You can specify custom data in this field.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("destination-country-parameters", "", "This field is used for any country-specific registration requirements.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("language-code", "", "Choose the language to use for the message.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("origination-identity", "", "The origination identity of the message.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("verification-channel", "", "Choose to send the verification code as an SMS or voice message.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.Flags().String("verified-destination-number-id", "", "The unique identifier for the verified destination phone number.")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.MarkFlagRequired("verification-channel")
	pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd.MarkFlagRequired("verified-destination-number-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_sendDestinationNumberVerificationCodeCmd)
}
