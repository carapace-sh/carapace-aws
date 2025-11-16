package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getEventStreamCmd = &cobra.Command{
	Use:   "get-event-stream",
	Short: "Retrieves information about the event stream settings for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getEventStreamCmd).Standalone()

	pinpoint_getEventStreamCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getEventStreamCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getEventStreamCmd)
}
