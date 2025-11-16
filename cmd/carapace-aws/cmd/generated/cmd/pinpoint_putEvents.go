package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_putEventsCmd = &cobra.Command{
	Use:   "put-events",
	Short: "Creates a new event to record for endpoints, or creates or updates endpoint data that existing events are associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_putEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_putEventsCmd).Standalone()

		pinpoint_putEventsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_putEventsCmd.Flags().String("events-request", "", "")
		pinpoint_putEventsCmd.MarkFlagRequired("application-id")
		pinpoint_putEventsCmd.MarkFlagRequired("events-request")
	})
	pinpointCmd.AddCommand(pinpoint_putEventsCmd)
}
