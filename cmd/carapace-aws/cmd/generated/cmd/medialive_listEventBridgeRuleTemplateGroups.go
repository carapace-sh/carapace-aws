package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listEventBridgeRuleTemplateGroupsCmd = &cobra.Command{
	Use:   "list-event-bridge-rule-template-groups",
	Short: "Lists eventbridge rule template groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listEventBridgeRuleTemplateGroupsCmd).Standalone()

	medialive_listEventBridgeRuleTemplateGroupsCmd.Flags().String("max-results", "", "")
	medialive_listEventBridgeRuleTemplateGroupsCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results in paginated list responses.")
	medialive_listEventBridgeRuleTemplateGroupsCmd.Flags().String("signal-map-identifier", "", "A signal map's identifier.")
	medialiveCmd.AddCommand(medialive_listEventBridgeRuleTemplateGroupsCmd)
}
