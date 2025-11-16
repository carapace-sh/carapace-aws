package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoice_getConfigurationSetEventDestinationsCmd = &cobra.Command{
	Use:   "get-configuration-set-event-destinations",
	Short: "Obtain information about an event destination, including the types of events it reports, the Amazon Resource Name (ARN) of the destination, and the name of the event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoice_getConfigurationSetEventDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoice_getConfigurationSetEventDestinationsCmd).Standalone()

		pinpointSmsVoice_getConfigurationSetEventDestinationsCmd.Flags().String("configuration-set-name", "", "ConfigurationSetName")
		pinpointSmsVoice_getConfigurationSetEventDestinationsCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointSmsVoiceCmd.AddCommand(pinpointSmsVoice_getConfigurationSetEventDestinationsCmd)
}
