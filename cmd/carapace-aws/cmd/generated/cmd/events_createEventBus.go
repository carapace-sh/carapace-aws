package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createEventBusCmd = &cobra.Command{
	Use:   "create-event-bus",
	Short: "Creates a new event bus within your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createEventBusCmd).Standalone()

	events_createEventBusCmd.Flags().String("dead-letter-config", "", "")
	events_createEventBusCmd.Flags().String("description", "", "The event bus description.")
	events_createEventBusCmd.Flags().String("event-source-name", "", "If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.")
	events_createEventBusCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on this event bus.")
	events_createEventBusCmd.Flags().String("log-config", "", "The logging configuration settings for the event bus.")
	events_createEventBusCmd.Flags().String("name", "", "The name of the new event bus.")
	events_createEventBusCmd.Flags().String("tags", "", "Tags to associate with the event bus.")
	events_createEventBusCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_createEventBusCmd)
}
