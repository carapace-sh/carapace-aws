package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updatePhoneNumberCmd = &cobra.Command{
	Use:   "update-phone-number",
	Short: "Updates the configuration of an existing origination phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updatePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_updatePhoneNumberCmd).Standalone()

		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("deletion-protection-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("international-sending-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("no-deletion-protection-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("no-international-sending-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("no-self-managed-opt-outs-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("no-two-way-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().String("opt-out-list-name", "", "The OptOutList to add the phone number to.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().String("phone-number-id", "", "The unique identifier of the phone number.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("self-managed-opt-outs-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().String("two-way-channel-arn", "", "The Amazon Resource Name (ARN) of the two way channel.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().String("two-way-channel-role", "", "An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flags().Bool("two-way-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flag("no-deletion-protection-enabled").Hidden = true
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flag("no-international-sending-enabled").Hidden = true
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flag("no-self-managed-opt-outs-enabled").Hidden = true
		pinpointSmsVoiceV2_updatePhoneNumberCmd.Flag("no-two-way-enabled").Hidden = true
		pinpointSmsVoiceV2_updatePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updatePhoneNumberCmd)
}
