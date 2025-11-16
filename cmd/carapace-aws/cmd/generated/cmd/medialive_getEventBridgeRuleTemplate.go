package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_getEventBridgeRuleTemplateCmd = &cobra.Command{
	Use:   "get-event-bridge-rule-template",
	Short: "Retrieves the specified eventbridge rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_getEventBridgeRuleTemplateCmd).Standalone()

	medialive_getEventBridgeRuleTemplateCmd.Flags().String("identifier", "", "An eventbridge rule template's identifier.")
	medialive_getEventBridgeRuleTemplateCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_getEventBridgeRuleTemplateCmd)
}
