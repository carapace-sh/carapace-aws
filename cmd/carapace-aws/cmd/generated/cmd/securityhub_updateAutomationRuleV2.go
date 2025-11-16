package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateAutomationRuleV2Cmd = &cobra.Command{
	Use:   "update-automation-rule-v2",
	Short: "Updates a V2 automation rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateAutomationRuleV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateAutomationRuleV2Cmd).Standalone()

		securityhub_updateAutomationRuleV2Cmd.Flags().String("actions", "", "A list of actions to be performed when the rule criteria is met.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("criteria", "", "The filtering type and configuration of the automation rule.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("description", "", "A description of the automation rule.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("identifier", "", "The ARN of the automation rule.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("rule-name", "", "The name of the automation rule.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("rule-order", "", "Represents a value for the rule priority.")
		securityhub_updateAutomationRuleV2Cmd.Flags().String("rule-status", "", "The status of the automation rule.")
		securityhub_updateAutomationRuleV2Cmd.MarkFlagRequired("identifier")
	})
	securityhubCmd.AddCommand(securityhub_updateAutomationRuleV2Cmd)
}
