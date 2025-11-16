package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_startTransactionCmd = &cobra.Command{
	Use:   "start-transaction",
	Short: "Starts a new transaction and returns its transaction ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_startTransactionCmd).Standalone()

	lakeformation_startTransactionCmd.Flags().String("transaction-type", "", "Indicates whether this transaction should be read only or read and write.")
	lakeformationCmd.AddCommand(lakeformation_startTransactionCmd)
}
