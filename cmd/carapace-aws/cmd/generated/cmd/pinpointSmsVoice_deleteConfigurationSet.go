package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_deleteConfigurationSetCmd = &cobra.Command{
	Use:   "delete-configuration-set",
	Short: "Deletes an existing configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_deleteConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoice_deleteConfigurationSetCmd).Standalone()

		pinpointSmsVoice_deleteConfigurationSetCmd.Flags().String("configuration-set-name", "", "ConfigurationSetName")
		pinpointSmsVoice_deleteConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_deleteConfigurationSetCmd)
}
