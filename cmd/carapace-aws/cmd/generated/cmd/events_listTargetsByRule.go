package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listTargetsByRuleCmd = &cobra.Command{
	Use:   "list-targets-by-rule",
	Short: "Lists the targets assigned to the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listTargetsByRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listTargetsByRuleCmd).Standalone()

		events_listTargetsByRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
		events_listTargetsByRuleCmd.Flags().String("limit", "", "The maximum number of results to return.")
		events_listTargetsByRuleCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
		events_listTargetsByRuleCmd.Flags().String("rule", "", "The name of the rule.")
		events_listTargetsByRuleCmd.MarkFlagRequired("rule")
	})
	eventsCmd.AddCommand(events_listTargetsByRuleCmd)
}
