package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_disableRuleCmd = &cobra.Command{
	Use:   "disable-rule",
	Short: "Disables the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_disableRuleCmd).Standalone()

	events_disableRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
	events_disableRuleCmd.Flags().String("name", "", "The name of the rule.")
	events_disableRuleCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_disableRuleCmd)
}
