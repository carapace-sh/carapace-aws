package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_releaseSenderIdCmd = &cobra.Command{
	Use:   "release-sender-id",
	Short: "Releases an existing sender ID in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_releaseSenderIdCmd).Standalone()

	pinpointSmsVoiceV2_releaseSenderIdCmd.Flags().String("iso-country-code", "", "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.")
	pinpointSmsVoiceV2_releaseSenderIdCmd.Flags().String("sender-id", "", "The sender ID to release.")
	pinpointSmsVoiceV2_releaseSenderIdCmd.MarkFlagRequired("iso-country-code")
	pinpointSmsVoiceV2_releaseSenderIdCmd.MarkFlagRequired("sender-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_releaseSenderIdCmd)
}
