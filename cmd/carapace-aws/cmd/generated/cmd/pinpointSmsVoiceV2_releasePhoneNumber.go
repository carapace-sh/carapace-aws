package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_releasePhoneNumberCmd = &cobra.Command{
	Use:   "release-phone-number",
	Short: "Releases an existing origination phone number in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_releasePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_releasePhoneNumberCmd).Standalone()

		pinpointSmsVoiceV2_releasePhoneNumberCmd.Flags().String("phone-number-id", "", "The PhoneNumberId or PhoneNumberArn of the phone number to release.")
		pinpointSmsVoiceV2_releasePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_releasePhoneNumberCmd)
}
