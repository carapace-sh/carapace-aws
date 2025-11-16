package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteEventCmd = &cobra.Command{
	Use:   "delete-event",
	Short: "Deletes the specified event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteEventCmd).Standalone()

		frauddetector_deleteEventCmd.Flags().String("delete-audit-history", "", "Specifies whether or not to delete any predictions associated with the event.")
		frauddetector_deleteEventCmd.Flags().String("event-id", "", "The ID of the event to delete.")
		frauddetector_deleteEventCmd.Flags().String("event-type-name", "", "The name of the event type.")
		frauddetector_deleteEventCmd.MarkFlagRequired("event-id")
		frauddetector_deleteEventCmd.MarkFlagRequired("event-type-name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteEventCmd)
}
