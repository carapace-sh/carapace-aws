package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_describeActiveReceiptRuleSetCmd = &cobra.Command{
	Use:   "describe-active-receipt-rule-set",
	Short: "Returns the metadata and receipt rules for the receipt rule set that is currently active.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_describeActiveReceiptRuleSetCmd).Standalone()

	sesCmd.AddCommand(ses_describeActiveReceiptRuleSetCmd)
}
