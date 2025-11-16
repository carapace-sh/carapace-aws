package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeEventBusCmd = &cobra.Command{
	Use:   "describe-event-bus",
	Short: "Displays details about an event bus in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeEventBusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeEventBusCmd).Standalone()

		events_describeEventBusCmd.Flags().String("name", "", "The name or ARN of the event bus to show details for.")
	})
	eventsCmd.AddCommand(events_describeEventBusCmd)
}
