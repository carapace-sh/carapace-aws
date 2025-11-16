package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteReceiptRuleSetCmd = &cobra.Command{
	Use:   "delete-receipt-rule-set",
	Short: "Deletes the specified receipt rule set and all of the receipt rules it contains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteReceiptRuleSetCmd).Standalone()

	ses_deleteReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set to delete.")
	ses_deleteReceiptRuleSetCmd.MarkFlagRequired("rule-set-name")
	sesCmd.AddCommand(ses_deleteReceiptRuleSetCmd)
}
