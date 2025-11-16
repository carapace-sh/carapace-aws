package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteAutomationRuleV2Cmd = &cobra.Command{
	Use:   "delete-automation-rule-v2",
	Short: "Deletes a V2 automation rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteAutomationRuleV2Cmd).Standalone()

	securityhub_deleteAutomationRuleV2Cmd.Flags().String("identifier", "", "The ARN of the V2 automation rule.")
	securityhub_deleteAutomationRuleV2Cmd.MarkFlagRequired("identifier")
	securityhubCmd.AddCommand(securityhub_deleteAutomationRuleV2Cmd)
}
