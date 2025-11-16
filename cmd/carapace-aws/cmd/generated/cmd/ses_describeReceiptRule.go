package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_describeReceiptRuleCmd = &cobra.Command{
	Use:   "describe-receipt-rule",
	Short: "Returns the details of the specified receipt rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_describeReceiptRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_describeReceiptRuleCmd).Standalone()

		ses_describeReceiptRuleCmd.Flags().String("rule-name", "", "The name of the receipt rule.")
		ses_describeReceiptRuleCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set that the receipt rule belongs to.")
		ses_describeReceiptRuleCmd.MarkFlagRequired("rule-name")
		ses_describeReceiptRuleCmd.MarkFlagRequired("rule-set-name")
	})
	sesCmd.AddCommand(ses_describeReceiptRuleCmd)
}
