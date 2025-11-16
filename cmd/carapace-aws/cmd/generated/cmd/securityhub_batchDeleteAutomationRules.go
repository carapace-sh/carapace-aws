package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchDeleteAutomationRulesCmd = &cobra.Command{
	Use:   "batch-delete-automation-rules",
	Short: "Deletes one or more automation rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchDeleteAutomationRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchDeleteAutomationRulesCmd).Standalone()

		securityhub_batchDeleteAutomationRulesCmd.Flags().String("automation-rules-arns", "", "A list of Amazon Resource Names (ARNs) for the rules that are to be deleted.")
		securityhub_batchDeleteAutomationRulesCmd.MarkFlagRequired("automation-rules-arns")
	})
	securityhubCmd.AddCommand(securityhub_batchDeleteAutomationRulesCmd)
}
