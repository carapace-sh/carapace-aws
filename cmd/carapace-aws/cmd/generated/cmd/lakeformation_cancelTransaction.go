package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_cancelTransactionCmd = &cobra.Command{
	Use:   "cancel-transaction",
	Short: "Attempts to cancel the specified transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_cancelTransactionCmd).Standalone()

	lakeformation_cancelTransactionCmd.Flags().String("transaction-id", "", "The transaction to cancel.")
	lakeformation_cancelTransactionCmd.MarkFlagRequired("transaction-id")
	lakeformationCmd.AddCommand(lakeformation_cancelTransactionCmd)
}
