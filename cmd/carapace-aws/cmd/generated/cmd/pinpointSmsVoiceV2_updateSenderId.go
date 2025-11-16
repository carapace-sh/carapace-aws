package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updateSenderIdCmd = &cobra.Command{
	Use:   "update-sender-id",
	Short: "Updates the configuration of an existing sender ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updateSenderIdCmd).Standalone()

	pinpointSmsVoiceV2_updateSenderIdCmd.Flags().Bool("deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updateSenderIdCmd.Flags().String("iso-country-code", "", "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.")
	pinpointSmsVoiceV2_updateSenderIdCmd.Flags().Bool("no-deletion-protection-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updateSenderIdCmd.Flags().String("sender-id", "", "The sender ID to update.")
	pinpointSmsVoiceV2_updateSenderIdCmd.MarkFlagRequired("iso-country-code")
	pinpointSmsVoiceV2_updateSenderIdCmd.Flag("no-deletion-protection-enabled").Hidden = true
	pinpointSmsVoiceV2_updateSenderIdCmd.MarkFlagRequired("sender-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updateSenderIdCmd)
}
