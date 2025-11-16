package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_putRuleCmd = &cobra.Command{
	Use:   "put-rule",
	Short: "Creates or updates the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_putRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_putRuleCmd).Standalone()

		events_putRuleCmd.Flags().String("description", "", "A description of the rule.")
		events_putRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus to associate with this rule.")
		events_putRuleCmd.Flags().String("event-pattern", "", "The event pattern.")
		events_putRuleCmd.Flags().String("name", "", "The name of the rule that you are creating or updating.")
		events_putRuleCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role associated with the rule.")
		events_putRuleCmd.Flags().String("schedule-expression", "", "The scheduling expression.")
		events_putRuleCmd.Flags().String("state", "", "The state of the rule.")
		events_putRuleCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the rule.")
		events_putRuleCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_putRuleCmd)
}
