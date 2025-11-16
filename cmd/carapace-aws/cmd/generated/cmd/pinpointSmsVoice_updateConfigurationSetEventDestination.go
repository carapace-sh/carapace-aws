package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_updateConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "update-configuration-set-event-destination",
	Short: "Update an event destination in a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_updateConfigurationSetEventDestinationCmd).Standalone()

	pinpointSmsVoice_updateConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "ConfigurationSetName")
	pinpointSmsVoice_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "")
	pinpointSmsVoice_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "EventDestinationName")
	pinpointSmsVoice_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoice_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_updateConfigurationSetEventDestinationCmd)
}
