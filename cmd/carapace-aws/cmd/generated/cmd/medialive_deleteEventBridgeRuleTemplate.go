package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteEventBridgeRuleTemplateCmd = &cobra.Command{
	Use:   "delete-event-bridge-rule-template",
	Short: "Deletes an eventbridge rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteEventBridgeRuleTemplateCmd).Standalone()

	medialive_deleteEventBridgeRuleTemplateCmd.Flags().String("identifier", "", "An eventbridge rule template's identifier.")
	medialive_deleteEventBridgeRuleTemplateCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_deleteEventBridgeRuleTemplateCmd)
}
