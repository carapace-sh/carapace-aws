package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getAutomationRuleV2Cmd = &cobra.Command{
	Use:   "get-automation-rule-v2",
	Short: "Returns an automation rule for the V2 service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getAutomationRuleV2Cmd).Standalone()

	securityhub_getAutomationRuleV2Cmd.Flags().String("identifier", "", "The ARN of the V2 automation rule.")
	securityhub_getAutomationRuleV2Cmd.MarkFlagRequired("identifier")
	securityhubCmd.AddCommand(securityhub_getAutomationRuleV2Cmd)
}
