package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createReceiptRuleSetCmd = &cobra.Command{
	Use:   "create-receipt-rule-set",
	Short: "Creates an empty receipt rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createReceiptRuleSetCmd).Standalone()

	ses_createReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the rule set to create.")
	ses_createReceiptRuleSetCmd.MarkFlagRequired("rule-set-name")
	sesCmd.AddCommand(ses_createReceiptRuleSetCmd)
}
