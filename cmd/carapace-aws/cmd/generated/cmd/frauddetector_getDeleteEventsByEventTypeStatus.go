package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getDeleteEventsByEventTypeStatusCmd = &cobra.Command{
	Use:   "get-delete-events-by-event-type-status",
	Short: "Retrieves the status of a `DeleteEventsByEventType` action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getDeleteEventsByEventTypeStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getDeleteEventsByEventTypeStatusCmd).Standalone()

		frauddetector_getDeleteEventsByEventTypeStatusCmd.Flags().String("event-type-name", "", "Name of event type for which to get the deletion status.")
		frauddetector_getDeleteEventsByEventTypeStatusCmd.MarkFlagRequired("event-type-name")
	})
	frauddetectorCmd.AddCommand(frauddetector_getDeleteEventsByEventTypeStatusCmd)
}
