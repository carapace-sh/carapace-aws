package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createReceiptRuleCmd = &cobra.Command{
	Use:   "create-receipt-rule",
	Short: "Creates a receipt rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createReceiptRuleCmd).Standalone()

	ses_createReceiptRuleCmd.Flags().String("after", "", "The name of an existing rule after which the new rule is placed.")
	ses_createReceiptRuleCmd.Flags().String("rule", "", "A data structure that contains the specified rule's name, actions, recipients, domains, enabled status, scan status, and TLS policy.")
	ses_createReceiptRuleCmd.Flags().String("rule-set-name", "", "The name of the rule set where the receipt rule is added.")
	ses_createReceiptRuleCmd.MarkFlagRequired("rule")
	ses_createReceiptRuleCmd.MarkFlagRequired("rule-set-name")
	sesCmd.AddCommand(ses_createReceiptRuleCmd)
}
