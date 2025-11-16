package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteEventsByEventTypeCmd = &cobra.Command{
	Use:   "delete-events-by-event-type",
	Short: "Deletes all events of a particular event type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteEventsByEventTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteEventsByEventTypeCmd).Standalone()

		frauddetector_deleteEventsByEventTypeCmd.Flags().String("event-type-name", "", "The name of the event type.")
		frauddetector_deleteEventsByEventTypeCmd.MarkFlagRequired("event-type-name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteEventsByEventTypeCmd)
}
