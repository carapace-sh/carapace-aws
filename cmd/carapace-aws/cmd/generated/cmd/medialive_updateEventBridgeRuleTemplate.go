package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateEventBridgeRuleTemplateCmd = &cobra.Command{
	Use:   "update-event-bridge-rule-template",
	Short: "Updates the specified eventbridge rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateEventBridgeRuleTemplateCmd).Standalone()

	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("description", "", "A resource's optional description.")
	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("event-targets", "", "")
	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("event-type", "", "")
	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("group-identifier", "", "An eventbridge rule template group's identifier.")
	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("identifier", "", "An eventbridge rule template's identifier.")
	medialive_updateEventBridgeRuleTemplateCmd.Flags().String("name", "", "A resource's name.")
	medialive_updateEventBridgeRuleTemplateCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_updateEventBridgeRuleTemplateCmd)
}
