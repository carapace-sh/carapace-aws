package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createEventBridgeRuleTemplateCmd = &cobra.Command{
	Use:   "create-event-bridge-rule-template",
	Short: "Creates an eventbridge rule template to monitor events and send notifications to your targeted resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createEventBridgeRuleTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createEventBridgeRuleTemplateCmd).Standalone()

		medialive_createEventBridgeRuleTemplateCmd.Flags().String("description", "", "A resource's optional description.")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("event-targets", "", "")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("event-type", "", "")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("group-identifier", "", "An eventbridge rule template group's identifier.")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("name", "", "A resource's name.")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
		medialive_createEventBridgeRuleTemplateCmd.Flags().String("tags", "", "")
		medialive_createEventBridgeRuleTemplateCmd.MarkFlagRequired("event-type")
		medialive_createEventBridgeRuleTemplateCmd.MarkFlagRequired("group-identifier")
		medialive_createEventBridgeRuleTemplateCmd.MarkFlagRequired("name")
	})
	medialiveCmd.AddCommand(medialive_createEventBridgeRuleTemplateCmd)
}
