package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteProtectConfigurationCmd = &cobra.Command{
	Use:   "delete-protect-configuration",
	Short: "Permanently delete the protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteProtectConfigurationCmd).Standalone()

	pinpointSmsVoiceV2_deleteProtectConfigurationCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_deleteProtectConfigurationCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteProtectConfigurationCmd)
}
