package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_putEventsCmd = &cobra.Command{
	Use:   "put-events",
	Short: "Sends custom events to Amazon EventBridge so that they can be matched to rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_putEventsCmd).Standalone()

	events_putEventsCmd.Flags().String("endpoint-id", "", "The URL subdomain of the endpoint.")
	events_putEventsCmd.Flags().String("entries", "", "The entry that defines an event in your system.")
	events_putEventsCmd.MarkFlagRequired("entries")
	eventsCmd.AddCommand(events_putEventsCmd)
}
