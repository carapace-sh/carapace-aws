package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_createConfigurationSetCmd = &cobra.Command{
	Use:   "create-configuration-set",
	Short: "Create a new configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_createConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoice_createConfigurationSetCmd).Standalone()

		pinpointSmsVoice_createConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name that you want to give the configuration set.")
	})
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_createConfigurationSetCmd)
}
