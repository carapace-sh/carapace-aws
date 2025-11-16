package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_requestPhoneNumberCmd = &cobra.Command{
	Use:   "request-phone-number",
	Short: "Request an origination phone number for use in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_requestPhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_requestPhoneNumberCmd).Standalone()

		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().Bool("deletion-protection-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().Bool("international-sending-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("iso-country-code", "", "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("message-type", "", "The type of message.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().Bool("no-deletion-protection-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().Bool("no-international-sending-enabled", false, "By default this is set to false.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("number-capabilities", "", "Indicates if the phone number will be used for text messages, voice messages, or both.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("number-type", "", "The type of phone number to request.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("opt-out-list-name", "", "The name of the OptOutList to associate with the phone number.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("pool-id", "", "The pool to associated with the phone number.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("registration-id", "", "Use this field to attach your phone number for an external registration process.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the requested phone number.")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.MarkFlagRequired("iso-country-code")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.MarkFlagRequired("message-type")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flag("no-deletion-protection-enabled").Hidden = true
		pinpointSmsVoiceV2_requestPhoneNumberCmd.Flag("no-international-sending-enabled").Hidden = true
		pinpointSmsVoiceV2_requestPhoneNumberCmd.MarkFlagRequired("number-capabilities")
		pinpointSmsVoiceV2_requestPhoneNumberCmd.MarkFlagRequired("number-type")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_requestPhoneNumberCmd)
}
