package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_createConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "create-configuration-set-event-destination",
	Short: "Create an event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_createConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_createConfigurationSetEventDestinationCmd).Standalone()

		pinpointEmail_createConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to add an event destination to.")
		pinpointEmail_createConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "An object that defines the event destination.")
		pinpointEmail_createConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "A name that identifies the event destination within the configuration set.")
		pinpointEmail_createConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		pinpointEmail_createConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
		pinpointEmail_createConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_createConfigurationSetEventDestinationCmd)
}
