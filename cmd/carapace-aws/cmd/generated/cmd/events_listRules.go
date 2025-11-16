package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "Lists your Amazon EventBridge rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listRulesCmd).Standalone()

	events_listRulesCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus to list the rules for.")
	events_listRulesCmd.Flags().String("limit", "", "The maximum number of results to return.")
	events_listRulesCmd.Flags().String("name-prefix", "", "The prefix matching the rule name.")
	events_listRulesCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	eventsCmd.AddCommand(events_listRulesCmd)
}
