package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "You must provide the name of the partition key attribute and a single value for that attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_queryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_queryCmd).Standalone()

		dynamodb_queryCmd.Flags().String("attributes-to-get", "", "This is a legacy parameter.")
		dynamodb_queryCmd.Flags().String("conditional-operator", "", "This is a legacy parameter.")
		dynamodb_queryCmd.Flags().String("consistent-read", "", "Determines the read consistency model: If set to `true`, then the operation uses strongly consistent reads; otherwise, the operation uses eventually consistent reads.")
		dynamodb_queryCmd.Flags().String("exclusive-start-key", "", "The primary key of the first item that this operation will evaluate.")
		dynamodb_queryCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_queryCmd.Flags().String("expression-attribute-values", "", "One or more values that can be substituted in an expression.")
		dynamodb_queryCmd.Flags().String("filter-expression", "", "A string that contains conditions that DynamoDB applies after the `Query` operation, but before the data is returned to you.")
		dynamodb_queryCmd.Flags().String("index-name", "", "The name of an index to query.")
		dynamodb_queryCmd.Flags().String("key-condition-expression", "", "The condition that specifies the key values for items to be retrieved by the `Query` action.")
		dynamodb_queryCmd.Flags().String("key-conditions", "", "This is a legacy parameter.")
		dynamodb_queryCmd.Flags().String("limit", "", "The maximum number of items to evaluate (not necessarily the number of matching items).")
		dynamodb_queryCmd.Flags().String("projection-expression", "", "A string that identifies one or more attributes to retrieve from the table.")
		dynamodb_queryCmd.Flags().String("query-filter", "", "This is a legacy parameter.")
		dynamodb_queryCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_queryCmd.Flags().String("scan-index-forward", "", "Specifies the order for index traversal: If `true` (default), the traversal is performed in ascending order; if `false`, the traversal is performed in descending order.")
		dynamodb_queryCmd.Flags().String("select", "", "The attributes to be returned in the result.")
		dynamodb_queryCmd.Flags().String("table-name", "", "The name of the table containing the requested items.")
		dynamodb_queryCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_queryCmd)
}
