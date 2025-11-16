package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_commitTransactionCmd = &cobra.Command{
	Use:   "commit-transaction",
	Short: "Attempts to commit the specified transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_commitTransactionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_commitTransactionCmd).Standalone()

		lakeformation_commitTransactionCmd.Flags().String("transaction-id", "", "The transaction to commit.")
		lakeformation_commitTransactionCmd.MarkFlagRequired("transaction-id")
	})
	lakeformationCmd.AddCommand(lakeformation_commitTransactionCmd)
}
