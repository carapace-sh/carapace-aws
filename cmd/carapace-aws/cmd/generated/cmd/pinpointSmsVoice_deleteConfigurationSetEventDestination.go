package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "delete-configuration-set-event-destination",
	Short: "Deletes an event destination in a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd).Standalone()

	pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "ConfigurationSetName")
	pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "EventDestinationName")
	pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_deleteConfigurationSetEventDestinationCmd)
}
