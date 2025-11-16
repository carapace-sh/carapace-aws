package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteEventBridgeRuleTemplateGroupCmd = &cobra.Command{
	Use:   "delete-event-bridge-rule-template-group",
	Short: "Deletes an eventbridge rule template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteEventBridgeRuleTemplateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteEventBridgeRuleTemplateGroupCmd).Standalone()

		medialive_deleteEventBridgeRuleTemplateGroupCmd.Flags().String("identifier", "", "An eventbridge rule template group's identifier.")
		medialive_deleteEventBridgeRuleTemplateGroupCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_deleteEventBridgeRuleTemplateGroupCmd)
}
