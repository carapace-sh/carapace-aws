package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_associateOriginationIdentityCmd = &cobra.Command{
	Use:   "associate-origination-identity",
	Short: "Associates the specified origination identity with a pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_associateOriginationIdentityCmd).Standalone()

	pinpointSmsVoiceV2_associateOriginationIdentityCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.Flags().String("iso-country-code", "", "The new two-character code, in ISO 3166-1 alpha-2 format, for the country or region of the origination identity.")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.Flags().String("origination-identity", "", "The origination identity to use, such as PhoneNumberId, PhoneNumberArn, SenderId, or SenderIdArn.")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.Flags().String("pool-id", "", "The pool to update with the new Identity.")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.MarkFlagRequired("iso-country-code")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.MarkFlagRequired("origination-identity")
	pinpointSmsVoiceV2_associateOriginationIdentityCmd.MarkFlagRequired("pool-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_associateOriginationIdentityCmd)
}
