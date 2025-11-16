package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createConfigurationSetCmd = &cobra.Command{
	Use:   "create-configuration-set",
	Short: "Creates a new configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createConfigurationSetCmd).Standalone()

	pinpointSmsVoiceV2_createConfigurationSetCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_createConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name to use for the new configuration set.")
	pinpointSmsVoiceV2_createConfigurationSetCmd.Flags().String("tags", "", "An array of key and value pair tags that's associated with the new configuration set.")
	pinpointSmsVoiceV2_createConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createConfigurationSetCmd)
}
