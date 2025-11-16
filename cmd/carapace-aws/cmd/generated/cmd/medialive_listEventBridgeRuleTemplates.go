package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listEventBridgeRuleTemplatesCmd = &cobra.Command{
	Use:   "list-event-bridge-rule-templates",
	Short: "Lists eventbridge rule templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listEventBridgeRuleTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listEventBridgeRuleTemplatesCmd).Standalone()

		medialive_listEventBridgeRuleTemplatesCmd.Flags().String("group-identifier", "", "An eventbridge rule template group's identifier.")
		medialive_listEventBridgeRuleTemplatesCmd.Flags().String("max-results", "", "")
		medialive_listEventBridgeRuleTemplatesCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results in paginated list responses.")
		medialive_listEventBridgeRuleTemplatesCmd.Flags().String("signal-map-identifier", "", "A signal map's identifier.")
	})
	medialiveCmd.AddCommand(medialive_listEventBridgeRuleTemplatesCmd)
}
