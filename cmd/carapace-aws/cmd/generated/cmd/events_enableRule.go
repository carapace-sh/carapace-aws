package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_enableRuleCmd = &cobra.Command{
	Use:   "enable-rule",
	Short: "Enables the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_enableRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_enableRuleCmd).Standalone()

		events_enableRuleCmd.Flags().String("event-bus-name", "", "The name or ARN of the event bus associated with the rule.")
		events_enableRuleCmd.Flags().String("name", "", "The name of the rule.")
		events_enableRuleCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_enableRuleCmd)
}
