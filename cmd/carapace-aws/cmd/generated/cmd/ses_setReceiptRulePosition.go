package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setReceiptRulePositionCmd = &cobra.Command{
	Use:   "set-receipt-rule-position",
	Short: "Sets the position of the specified receipt rule in the receipt rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setReceiptRulePositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_setReceiptRulePositionCmd).Standalone()

		ses_setReceiptRulePositionCmd.Flags().String("after", "", "The name of the receipt rule after which to place the specified receipt rule.")
		ses_setReceiptRulePositionCmd.Flags().String("rule-name", "", "The name of the receipt rule to reposition.")
		ses_setReceiptRulePositionCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set that contains the receipt rule to reposition.")
		ses_setReceiptRulePositionCmd.MarkFlagRequired("rule-name")
		ses_setReceiptRulePositionCmd.MarkFlagRequired("rule-set-name")
	})
	sesCmd.AddCommand(ses_setReceiptRulePositionCmd)
}
