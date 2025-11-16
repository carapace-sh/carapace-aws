package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_transactWriteItemsCmd = &cobra.Command{
	Use:   "transact-write-items",
	Short: "`TransactWriteItems` is a synchronous write operation that groups up to 100 action requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_transactWriteItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_transactWriteItemsCmd).Standalone()

		dynamodb_transactWriteItemsCmd.Flags().String("client-request-token", "", "Providing a `ClientRequestToken` makes the call to `TransactWriteItems` idempotent, meaning that multiple identical calls have the same effect as one single call.")
		dynamodb_transactWriteItemsCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_transactWriteItemsCmd.Flags().String("return-item-collection-metrics", "", "Determines whether item collection metrics are returned.")
		dynamodb_transactWriteItemsCmd.Flags().String("transact-items", "", "An ordered array of up to 100 `TransactWriteItem` objects, each of which contains a `ConditionCheck`, `Put`, `Update`, or `Delete` object.")
		dynamodb_transactWriteItemsCmd.MarkFlagRequired("transact-items")
	})
	dynamodbCmd.AddCommand(dynamodb_transactWriteItemsCmd)
}
