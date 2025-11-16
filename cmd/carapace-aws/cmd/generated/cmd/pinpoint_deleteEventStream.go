package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteEventStreamCmd = &cobra.Command{
	Use:   "delete-event-stream",
	Short: "Deletes the event stream for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteEventStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteEventStreamCmd).Standalone()

		pinpoint_deleteEventStreamCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteEventStreamCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteEventStreamCmd)
}
