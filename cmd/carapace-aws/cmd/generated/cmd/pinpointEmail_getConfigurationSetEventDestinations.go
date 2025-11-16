package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getConfigurationSetEventDestinationsCmd = &cobra.Command{
	Use:   "get-configuration-set-event-destinations",
	Short: "Retrieve a list of event destinations that are associated with a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getConfigurationSetEventDestinationsCmd).Standalone()

	pinpointEmail_getConfigurationSetEventDestinationsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination.")
	pinpointEmail_getConfigurationSetEventDestinationsCmd.MarkFlagRequired("configuration-set-name")
	pinpointEmailCmd.AddCommand(pinpointEmail_getConfigurationSetEventDestinationsCmd)
}
