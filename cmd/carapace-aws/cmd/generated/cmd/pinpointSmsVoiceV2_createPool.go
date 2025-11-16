package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createPoolCmd = &cobra.Command{
	Use:   "create-pool",
	Short: "Creates a new pool and associates the specified origination identity to the pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createPoolCmd).Standalone()

	pinpointSmsVoiceV2_createPoolCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().Bool("deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().String("iso-country-code", "", "The new two-character code, in ISO 3166-1 alpha-2 format, for the country or region of the new pool.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().String("message-type", "", "The type of message.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().Bool("no-deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().String("origination-identity", "", "The origination identity to use such as a PhoneNumberId, PhoneNumberArn, SenderId or SenderIdArn.")
	pinpointSmsVoiceV2_createPoolCmd.Flags().String("tags", "", "An array of tags (key and value pairs) associated with the pool.")
	pinpointSmsVoiceV2_createPoolCmd.MarkFlagRequired("iso-country-code")
	pinpointSmsVoiceV2_createPoolCmd.MarkFlagRequired("message-type")
	pinpointSmsVoiceV2_createPoolCmd.Flag("no-deletion-protection-enabled").Hidden = true
	pinpointSmsVoiceV2_createPoolCmd.MarkFlagRequired("origination-identity")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createPoolCmd)
}
