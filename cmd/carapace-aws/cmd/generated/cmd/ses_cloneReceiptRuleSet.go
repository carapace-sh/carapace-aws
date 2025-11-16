package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_cloneReceiptRuleSetCmd = &cobra.Command{
	Use:   "clone-receipt-rule-set",
	Short: "Creates a receipt rule set by cloning an existing one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_cloneReceiptRuleSetCmd).Standalone()

	ses_cloneReceiptRuleSetCmd.Flags().String("original-rule-set-name", "", "The name of the rule set to clone.")
	ses_cloneReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the rule set to create.")
	ses_cloneReceiptRuleSetCmd.MarkFlagRequired("original-rule-set-name")
	ses_cloneReceiptRuleSetCmd.MarkFlagRequired("rule-set-name")
	sesCmd.AddCommand(ses_cloneReceiptRuleSetCmd)
}
