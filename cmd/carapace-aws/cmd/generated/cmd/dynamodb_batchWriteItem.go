package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_batchWriteItemCmd = &cobra.Command{
	Use:   "batch-write-item",
	Short: "The `BatchWriteItem` operation puts or deletes multiple items in one or more tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_batchWriteItemCmd).Standalone()

	dynamodb_batchWriteItemCmd.Flags().String("request-items", "", "A map of one or more table names or table ARNs and, for each table, a list of operations to be performed (`DeleteRequest` or `PutRequest`).")
	dynamodb_batchWriteItemCmd.Flags().String("return-consumed-capacity", "", "")
	dynamodb_batchWriteItemCmd.Flags().String("return-item-collection-metrics", "", "Determines whether item collection metrics are returned.")
	dynamodb_batchWriteItemCmd.MarkFlagRequired("request-items")
	dynamodbCmd.AddCommand(dynamodb_batchWriteItemCmd)
}
