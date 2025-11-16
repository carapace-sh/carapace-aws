package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteEventBusCmd = &cobra.Command{
	Use:   "delete-event-bus",
	Short: "Deletes the specified custom event bus or partner event bus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteEventBusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_deleteEventBusCmd).Standalone()

		events_deleteEventBusCmd.Flags().String("name", "", "The name of the event bus to delete.")
		events_deleteEventBusCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_deleteEventBusCmd)
}
