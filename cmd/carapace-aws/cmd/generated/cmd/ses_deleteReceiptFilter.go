package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteReceiptFilterCmd = &cobra.Command{
	Use:   "delete-receipt-filter",
	Short: "Deletes the specified IP address filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteReceiptFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_deleteReceiptFilterCmd).Standalone()

		ses_deleteReceiptFilterCmd.Flags().String("filter-name", "", "The name of the IP address filter to delete.")
		ses_deleteReceiptFilterCmd.MarkFlagRequired("filter-name")
	})
	sesCmd.AddCommand(ses_deleteReceiptFilterCmd)
}
