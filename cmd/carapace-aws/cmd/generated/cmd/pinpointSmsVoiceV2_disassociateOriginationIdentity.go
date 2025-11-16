package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_disassociateOriginationIdentityCmd = &cobra.Command{
	Use:   "disassociate-origination-identity",
	Short: "Removes the specified origination identity from an existing pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_disassociateOriginationIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_disassociateOriginationIdentityCmd).Standalone()

		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.Flags().String("iso-country-code", "", "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.Flags().String("origination-identity", "", "The origination identity to use such as a PhoneNumberId, PhoneNumberArn, SenderId or SenderIdArn.")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.Flags().String("pool-id", "", "The unique identifier for the pool to disassociate with the origination identity.")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.MarkFlagRequired("iso-country-code")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.MarkFlagRequired("origination-identity")
		pinpointSmsVoiceV2_disassociateOriginationIdentityCmd.MarkFlagRequired("pool-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_disassociateOriginationIdentityCmd)
}
