package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listReceiptFiltersCmd = &cobra.Command{
	Use:   "list-receipt-filters",
	Short: "Lists the IP address filters associated with your Amazon Web Services account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listReceiptFiltersCmd).Standalone()

	sesCmd.AddCommand(ses_listReceiptFiltersCmd)
}
