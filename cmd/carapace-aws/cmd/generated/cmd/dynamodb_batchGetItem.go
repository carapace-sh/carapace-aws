package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_batchGetItemCmd = &cobra.Command{
	Use:   "batch-get-item",
	Short: "The `BatchGetItem` operation returns the attributes of one or more items from one or more tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_batchGetItemCmd).Standalone()

	dynamodb_batchGetItemCmd.Flags().String("request-items", "", "A map of one or more table names or table ARNs and, for each table, a map that describes one or more items to retrieve from that table.")
	dynamodb_batchGetItemCmd.Flags().String("return-consumed-capacity", "", "")
	dynamodb_batchGetItemCmd.MarkFlagRequired("request-items")
	dynamodbCmd.AddCommand(dynamodb_batchGetItemCmd)
}
