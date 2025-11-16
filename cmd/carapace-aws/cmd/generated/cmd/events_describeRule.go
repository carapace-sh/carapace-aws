package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeRuleCmd = &cobra.Command{
	Use:   "describe-rule",
	Short: "Describes the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeRuleCmd).Standalone()

		events_describeRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
		events_describeRuleCmd.Flags().String("name", "", "The name of the rule.")
		events_describeRuleCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_describeRuleCmd)
}
