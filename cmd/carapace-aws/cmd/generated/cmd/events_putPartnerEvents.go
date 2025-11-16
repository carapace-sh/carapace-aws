package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_putPartnerEventsCmd = &cobra.Command{
	Use:   "put-partner-events",
	Short: "This is used by SaaS partners to write events to a customer's partner event bus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_putPartnerEventsCmd).Standalone()

	events_putPartnerEventsCmd.Flags().String("entries", "", "The list of events to write to the event bus.")
	events_putPartnerEventsCmd.MarkFlagRequired("entries")
	eventsCmd.AddCommand(events_putPartnerEventsCmd)
}
