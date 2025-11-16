package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_describeTransactionCmd = &cobra.Command{
	Use:   "describe-transaction",
	Short: "Returns the details of a single transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_describeTransactionCmd).Standalone()

	lakeformation_describeTransactionCmd.Flags().String("transaction-id", "", "The transaction for which to return status.")
	lakeformation_describeTransactionCmd.MarkFlagRequired("transaction-id")
	lakeformationCmd.AddCommand(lakeformation_describeTransactionCmd)
}
