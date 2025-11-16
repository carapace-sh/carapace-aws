package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createAutomationRuleCmd = &cobra.Command{
	Use:   "create-automation-rule",
	Short: "Creates an automation rule based on input parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createAutomationRuleCmd).Standalone()

	securityhub_createAutomationRuleCmd.Flags().String("actions", "", "One or more actions to update finding fields if a finding matches the conditions specified in `Criteria`.")
	securityhub_createAutomationRuleCmd.Flags().String("criteria", "", "A set of ASFF finding field attributes and corresponding expected values that Security Hub uses to filter findings.")
	securityhub_createAutomationRuleCmd.Flags().String("description", "", "A description of the rule.")
	securityhub_createAutomationRuleCmd.Flags().Bool("is-terminal", false, "Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria.")
	securityhub_createAutomationRuleCmd.Flags().Bool("no-is-terminal", false, "Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria.")
	securityhub_createAutomationRuleCmd.Flags().String("rule-name", "", "The name of the rule.")
	securityhub_createAutomationRuleCmd.Flags().String("rule-order", "", "An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings.")
	securityhub_createAutomationRuleCmd.Flags().String("rule-status", "", "Whether the rule is active after it is created.")
	securityhub_createAutomationRuleCmd.Flags().String("tags", "", "User-defined tags associated with an automation rule.")
	securityhub_createAutomationRuleCmd.MarkFlagRequired("actions")
	securityhub_createAutomationRuleCmd.MarkFlagRequired("criteria")
	securityhub_createAutomationRuleCmd.MarkFlagRequired("description")
	securityhub_createAutomationRuleCmd.Flag("no-is-terminal").Hidden = true
	securityhub_createAutomationRuleCmd.MarkFlagRequired("rule-name")
	securityhub_createAutomationRuleCmd.MarkFlagRequired("rule-order")
	securityhubCmd.AddCommand(securityhub_createAutomationRuleCmd)
}
