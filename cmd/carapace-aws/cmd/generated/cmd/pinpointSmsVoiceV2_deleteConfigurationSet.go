package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteConfigurationSetCmd = &cobra.Command{
	Use:   "delete-configuration-set",
	Short: "Deletes an existing configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteConfigurationSetCmd).Standalone()

		pinpointSmsVoiceV2_deleteConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set or the configuration set ARN that you want to delete.")
		pinpointSmsVoiceV2_deleteConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteConfigurationSetCmd)
}
