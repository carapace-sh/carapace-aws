package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateReceiptRuleCmd = &cobra.Command{
	Use:   "update-receipt-rule",
	Short: "Updates a receipt rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateReceiptRuleCmd).Standalone()

	ses_updateReceiptRuleCmd.Flags().String("rule", "", "A data structure that contains the updated receipt rule information.")
	ses_updateReceiptRuleCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set that the receipt rule belongs to.")
	ses_updateReceiptRuleCmd.MarkFlagRequired("rule")
	ses_updateReceiptRuleCmd.MarkFlagRequired("rule-set-name")
	sesCmd.AddCommand(ses_updateReceiptRuleCmd)
}
