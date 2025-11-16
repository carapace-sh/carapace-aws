package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setActiveReceiptRuleSetCmd = &cobra.Command{
	Use:   "set-active-receipt-rule-set",
	Short: "Sets the specified receipt rule set as the active receipt rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setActiveReceiptRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_setActiveReceiptRuleSetCmd).Standalone()

		ses_setActiveReceiptRuleSetCmd.Flags().String("rule-set-name", "", "The name of the receipt rule set to make active.")
	})
	sesCmd.AddCommand(ses_setActiveReceiptRuleSetCmd)
}
