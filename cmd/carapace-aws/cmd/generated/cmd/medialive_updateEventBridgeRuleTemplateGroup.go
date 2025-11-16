package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateEventBridgeRuleTemplateGroupCmd = &cobra.Command{
	Use:   "update-event-bridge-rule-template-group",
	Short: "Updates the specified eventbridge rule template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateEventBridgeRuleTemplateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateEventBridgeRuleTemplateGroupCmd).Standalone()

		medialive_updateEventBridgeRuleTemplateGroupCmd.Flags().String("description", "", "A resource's optional description.")
		medialive_updateEventBridgeRuleTemplateGroupCmd.Flags().String("identifier", "", "An eventbridge rule template group's identifier.")
		medialive_updateEventBridgeRuleTemplateGroupCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_updateEventBridgeRuleTemplateGroupCmd)
}
