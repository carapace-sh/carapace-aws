package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_updateEventBusCmd = &cobra.Command{
	Use:   "update-event-bus",
	Short: "Updates the specified event bus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_updateEventBusCmd).Standalone()

	events_updateEventBusCmd.Flags().String("dead-letter-config", "", "")
	events_updateEventBusCmd.Flags().String("description", "", "The event bus description.")
	events_updateEventBusCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on this event bus.")
	events_updateEventBusCmd.Flags().String("log-config", "", "The logging configuration settings for the event bus.")
	events_updateEventBusCmd.Flags().String("name", "", "The name of the event bus.")
	eventsCmd.AddCommand(events_updateEventBusCmd)
}
