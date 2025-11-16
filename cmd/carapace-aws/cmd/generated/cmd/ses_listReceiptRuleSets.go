package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listReceiptRuleSetsCmd = &cobra.Command{
	Use:   "list-receipt-rule-sets",
	Short: "Lists the receipt rule sets that exist under your Amazon Web Services account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listReceiptRuleSetsCmd).Standalone()

	ses_listReceiptRuleSetsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListReceiptRuleSets` to indicate the position in the receipt rule set list.")
	sesCmd.AddCommand(ses_listReceiptRuleSetsCmd)
}
