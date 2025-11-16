package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_requestSenderIdCmd = &cobra.Command{
	Use:   "request-sender-id",
	Short: "Request a new sender ID that doesn't require registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_requestSenderIdCmd).Standalone()

	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().Bool("deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().String("iso-country-code", "", "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().String("message-types", "", "The type of message.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().Bool("no-deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().String("sender-id", "", "The sender ID string to request.")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the sender ID.")
	pinpointSmsVoiceV2_requestSenderIdCmd.MarkFlagRequired("iso-country-code")
	pinpointSmsVoiceV2_requestSenderIdCmd.Flag("no-deletion-protection-enabled").Hidden = true
	pinpointSmsVoiceV2_requestSenderIdCmd.MarkFlagRequired("sender-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_requestSenderIdCmd)
}
