package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_describeReceiptRuleSetCmd = &cobra.Command{
	Use:   "describe-receipt-rule-set",
	Short: "Returns the details of the specified receipt rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_describeReceiptRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_describeReceiptRuleSetCmd).Standalone()

		ses_describeReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set to describe.")
		ses_describeReceiptRuleSetCmd.MarkFlagRequired("rule-set-name")
	})
	sesCmd.AddCommand(ses_describeReceiptRuleSetCmd)
}
