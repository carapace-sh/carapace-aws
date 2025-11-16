package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deactivateEventSourceCmd = &cobra.Command{
	Use:   "deactivate-event-source",
	Short: "You can use this operation to temporarily stop receiving events from the specified partner event source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deactivateEventSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_deactivateEventSourceCmd).Standalone()

		events_deactivateEventSourceCmd.Flags().String("name", "", "The name of the partner event source to deactivate.")
		events_deactivateEventSourceCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_deactivateEventSourceCmd)
}
