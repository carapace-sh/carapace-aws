package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_extendTransactionCmd = &cobra.Command{
	Use:   "extend-transaction",
	Short: "Indicates to the service that the specified transaction is still active and should not be treated as idle and aborted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_extendTransactionCmd).Standalone()

	lakeformation_extendTransactionCmd.Flags().String("transaction-id", "", "The transaction to extend.")
	lakeformationCmd.AddCommand(lakeformation_extendTransactionCmd)
}
