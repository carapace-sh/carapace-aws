package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_updateConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "update-configuration-set-event-destination",
	Short: "Update the configuration of an event destination for a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_updateConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_updateConfigurationSetEventDestinationCmd).Standalone()

		pinpointEmail_updateConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination that you want to modify.")
		pinpointEmail_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "An object that defines the event destination.")
		pinpointEmail_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination that you want to modify.")
		pinpointEmail_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		pinpointEmail_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
		pinpointEmail_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_updateConfigurationSetEventDestinationCmd)
}
