package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_deleteConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "delete-configuration-set-event-destination",
	Short: "Delete an event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_deleteConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_deleteConfigurationSetEventDestinationCmd).Standalone()

		pinpointEmail_deleteConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination that you want to delete.")
		pinpointEmail_deleteConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination that you want to delete.")
		pinpointEmail_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		pinpointEmail_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_deleteConfigurationSetEventDestinationCmd)
}
