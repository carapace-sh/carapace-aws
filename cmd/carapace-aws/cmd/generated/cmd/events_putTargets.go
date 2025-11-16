package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_putTargetsCmd = &cobra.Command{
	Use:   "put-targets",
	Short: "Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_putTargetsCmd).Standalone()

	events_putTargetsCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
	events_putTargetsCmd.Flags().String("rule", "", "The name of the rule.")
	events_putTargetsCmd.Flags().String("targets", "", "The targets to update or add to the rule.")
	events_putTargetsCmd.MarkFlagRequired("rule")
	events_putTargetsCmd.MarkFlagRequired("targets")
	eventsCmd.AddCommand(events_putTargetsCmd)
}
