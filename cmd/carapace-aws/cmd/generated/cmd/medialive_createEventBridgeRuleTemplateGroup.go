package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createEventBridgeRuleTemplateGroupCmd = &cobra.Command{
	Use:   "create-event-bridge-rule-template-group",
	Short: "Creates an eventbridge rule template group to group your eventbridge rule templates and to attach to signal maps for dynamically creating notification rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createEventBridgeRuleTemplateGroupCmd).Standalone()

	medialive_createEventBridgeRuleTemplateGroupCmd.Flags().String("description", "", "A resource's optional description.")
	medialive_createEventBridgeRuleTemplateGroupCmd.Flags().String("name", "", "A resource's name.")
	medialive_createEventBridgeRuleTemplateGroupCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
	medialive_createEventBridgeRuleTemplateGroupCmd.Flags().String("tags", "", "")
	medialive_createEventBridgeRuleTemplateGroupCmd.MarkFlagRequired("name")
	medialiveCmd.AddCommand(medialive_createEventBridgeRuleTemplateGroupCmd)
}
