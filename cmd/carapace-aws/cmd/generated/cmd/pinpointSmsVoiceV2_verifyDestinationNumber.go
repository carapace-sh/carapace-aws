package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_verifyDestinationNumberCmd = &cobra.Command{
	Use:   "verify-destination-number",
	Short: "Use the verification code that was received by the verified destination phone number to opt-in the verified destination phone number to receive more messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_verifyDestinationNumberCmd).Standalone()

	pinpointSmsVoiceV2_verifyDestinationNumberCmd.Flags().String("verification-code", "", "The verification code that was received by the verified destination phone number.")
	pinpointSmsVoiceV2_verifyDestinationNumberCmd.Flags().String("verified-destination-number-id", "", "The unique identifier for the verififed destination phone number.")
	pinpointSmsVoiceV2_verifyDestinationNumberCmd.MarkFlagRequired("verification-code")
	pinpointSmsVoiceV2_verifyDestinationNumberCmd.MarkFlagRequired("verified-destination-number-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_verifyDestinationNumberCmd)
}
