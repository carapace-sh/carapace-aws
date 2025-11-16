package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteEventDestinationCmd = &cobra.Command{
	Use:   "delete-event-destination",
	Short: "Deletes an existing event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteEventDestinationCmd).Standalone()

	pinpointSmsVoiceV2_deleteEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set or the configuration set's Amazon Resource Name (ARN) to remove the event destination from.")
	pinpointSmsVoiceV2_deleteEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination to delete.")
	pinpointSmsVoiceV2_deleteEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2_deleteEventDestinationCmd.MarkFlagRequired("event-destination-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteEventDestinationCmd)
}
