package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_createConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "create-configuration-set-event-destination",
	Short: "Create a new event destination in a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_createConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoice_createConfigurationSetEventDestinationCmd).Standalone()

		pinpointSmsVoice_createConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "ConfigurationSetName")
		pinpointSmsVoice_createConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "")
		pinpointSmsVoice_createConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "A name that identifies the event destination.")
		pinpointSmsVoice_createConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_createConfigurationSetEventDestinationCmd)
}
