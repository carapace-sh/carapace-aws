package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createAutomationRuleV2Cmd = &cobra.Command{
	Use:   "create-automation-rule-v2",
	Short: "Creates a V2 automation rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createAutomationRuleV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createAutomationRuleV2Cmd).Standalone()

		securityhub_createAutomationRuleV2Cmd.Flags().String("actions", "", "A list of actions to be performed when the rule criteria is met.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("client-token", "", "A unique identifier used to ensure idempotency.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("criteria", "", "The filtering type and configuration of the automation rule.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("description", "", "A description of the V2 automation rule.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("rule-name", "", "The name of the V2 automation rule.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("rule-order", "", "The value for the rule priority.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("rule-status", "", "The status of the V2 automation rule.")
		securityhub_createAutomationRuleV2Cmd.Flags().String("tags", "", "A list of key-value pairs associated with the V2 automation rule.")
		securityhub_createAutomationRuleV2Cmd.MarkFlagRequired("actions")
		securityhub_createAutomationRuleV2Cmd.MarkFlagRequired("criteria")
		securityhub_createAutomationRuleV2Cmd.MarkFlagRequired("description")
		securityhub_createAutomationRuleV2Cmd.MarkFlagRequired("rule-name")
		securityhub_createAutomationRuleV2Cmd.MarkFlagRequired("rule-order")
	})
	securityhubCmd.AddCommand(securityhub_createAutomationRuleV2Cmd)
}
