package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_reorderReceiptRuleSetCmd = &cobra.Command{
	Use:   "reorder-receipt-rule-set",
	Short: "Reorders the receipt rules within a receipt rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_reorderReceiptRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_reorderReceiptRuleSetCmd).Standalone()

		ses_reorderReceiptRuleSetCmd.Flags().String("rule-names", "", "The specified receipt rule set's receipt rules, in order.")
		ses_reorderReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set to reorder.")
		ses_reorderReceiptRuleSetCmd.MarkFlagRequired("rule-names")
		ses_reorderReceiptRuleSetCmd.MarkFlagRequired("rule-set-name")
	})
	sesCmd.AddCommand(ses_reorderReceiptRuleSetCmd)
}
