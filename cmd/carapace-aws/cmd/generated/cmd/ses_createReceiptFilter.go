package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createReceiptFilterCmd = &cobra.Command{
	Use:   "create-receipt-filter",
	Short: "Creates a new IP address filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createReceiptFilterCmd).Standalone()

	ses_createReceiptFilterCmd.Flags().String("filter", "", "A data structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.")
	ses_createReceiptFilterCmd.MarkFlagRequired("filter")
	sesCmd.AddCommand(ses_createReceiptFilterCmd)
}
