package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteEventTypeCmd = &cobra.Command{
	Use:   "delete-event-type",
	Short: "Deletes an event type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteEventTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteEventTypeCmd).Standalone()

		frauddetector_deleteEventTypeCmd.Flags().String("name", "", "The name of the event type to delete.")
		frauddetector_deleteEventTypeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteEventTypeCmd)
}
