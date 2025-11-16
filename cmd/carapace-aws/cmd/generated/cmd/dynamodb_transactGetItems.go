package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_transactGetItemsCmd = &cobra.Command{
	Use:   "transact-get-items",
	Short: "`TransactGetItems` is a synchronous operation that atomically retrieves multiple items from one or more tables (but not from indexes) in a single account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_transactGetItemsCmd).Standalone()

	dynamodb_transactGetItemsCmd.Flags().String("return-consumed-capacity", "", "A value of `TOTAL` causes consumed capacity information to be returned, and a value of `NONE` prevents that information from being returned.")
	dynamodb_transactGetItemsCmd.Flags().String("transact-items", "", "An ordered array of up to 100 `TransactGetItem` objects, each of which contains a `Get` structure.")
	dynamodb_transactGetItemsCmd.MarkFlagRequired("transact-items")
	dynamodbCmd.AddCommand(dynamodb_transactGetItemsCmd)
}
