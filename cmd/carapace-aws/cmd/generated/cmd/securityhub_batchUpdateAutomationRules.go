package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchUpdateAutomationRulesCmd = &cobra.Command{
	Use:   "batch-update-automation-rules",
	Short: "Updates one or more automation rules based on rule Amazon Resource Names (ARNs) and input parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchUpdateAutomationRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchUpdateAutomationRulesCmd).Standalone()

		securityhub_batchUpdateAutomationRulesCmd.Flags().String("update-automation-rules-request-items", "", "An array of ARNs for the rules that are to be updated.")
		securityhub_batchUpdateAutomationRulesCmd.MarkFlagRequired("update-automation-rules-request-items")
	})
	securityhubCmd.AddCommand(securityhub_batchUpdateAutomationRulesCmd)
}
