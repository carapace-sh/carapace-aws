package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchGetAutomationRulesCmd = &cobra.Command{
	Use:   "batch-get-automation-rules",
	Short: "Retrieves a list of details for automation rules based on rule Amazon Resource Names (ARNs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchGetAutomationRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchGetAutomationRulesCmd).Standalone()

		securityhub_batchGetAutomationRulesCmd.Flags().String("automation-rules-arns", "", "A list of rule ARNs to get details for.")
		securityhub_batchGetAutomationRulesCmd.MarkFlagRequired("automation-rules-arns")
	})
	securityhubCmd.AddCommand(securityhub_batchGetAutomationRulesCmd)
}
