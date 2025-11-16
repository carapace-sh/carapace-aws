package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_activateEventSourceCmd = &cobra.Command{
	Use:   "activate-event-source",
	Short: "Activates a partner event source that has been deactivated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_activateEventSourceCmd).Standalone()

	events_activateEventSourceCmd.Flags().String("name", "", "The name of the partner event source to activate.")
	events_activateEventSourceCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_activateEventSourceCmd)
}
