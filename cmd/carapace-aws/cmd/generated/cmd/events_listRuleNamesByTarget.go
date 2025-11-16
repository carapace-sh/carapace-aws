package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listRuleNamesByTargetCmd = &cobra.Command{
	Use:   "list-rule-names-by-target",
	Short: "Lists the rules for the specified target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listRuleNamesByTargetCmd).Standalone()

	events_listRuleNamesByTargetCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus to list rules for.")
	events_listRuleNamesByTargetCmd.Flags().String("limit", "", "The maximum number of results to return.")
	events_listRuleNamesByTargetCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	events_listRuleNamesByTargetCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the target resource.")
	events_listRuleNamesByTargetCmd.MarkFlagRequired("target-arn")
	eventsCmd.AddCommand(events_listRuleNamesByTargetCmd)
}
