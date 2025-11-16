package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteReceiptRuleCmd = &cobra.Command{
	Use:   "delete-receipt-rule",
	Short: "Deletes the specified receipt rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteReceiptRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_deleteReceiptRuleCmd).Standalone()

		ses_deleteReceiptRuleCmd.Flags().String("rule-name", "", "The name of the receipt rule to delete.")
		ses_deleteReceiptRuleCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set that contains the receipt rule to delete.")
		ses_deleteReceiptRuleCmd.MarkFlagRequired("rule-name")
		ses_deleteReceiptRuleCmd.MarkFlagRequired("rule-set-name")
	})
	sesCmd.AddCommand(ses_deleteReceiptRuleCmd)
}
