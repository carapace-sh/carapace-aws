package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_getEventBridgeRuleTemplateGroupCmd = &cobra.Command{
	Use:   "get-event-bridge-rule-template-group",
	Short: "Retrieves the specified eventbridge rule template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_getEventBridgeRuleTemplateGroupCmd).Standalone()

	medialive_getEventBridgeRuleTemplateGroupCmd.Flags().String("identifier", "", "An eventbridge rule template group's identifier.")
	medialive_getEventBridgeRuleTemplateGroupCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_getEventBridgeRuleTemplateGroupCmd)
}
