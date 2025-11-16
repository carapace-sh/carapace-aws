package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_deleteItemCmd = &cobra.Command{
	Use:   "delete-item",
	Short: "Deletes a single item in a table by primary key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_deleteItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_deleteItemCmd).Standalone()

		dynamodb_deleteItemCmd.Flags().String("condition-expression", "", "A condition that must be satisfied in order for a conditional `DeleteItem` to succeed.")
		dynamodb_deleteItemCmd.Flags().String("conditional-operator", "", "This is a legacy parameter.")
		dynamodb_deleteItemCmd.Flags().String("expected", "", "This is a legacy parameter.")
		dynamodb_deleteItemCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_deleteItemCmd.Flags().String("expression-attribute-values", "", "One or more values that can be substituted in an expression.")
		dynamodb_deleteItemCmd.Flags().String("key", "", "A map of attribute names to `AttributeValue` objects, representing the primary key of the item to delete.")
		dynamodb_deleteItemCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_deleteItemCmd.Flags().String("return-item-collection-metrics", "", "Determines whether item collection metrics are returned.")
		dynamodb_deleteItemCmd.Flags().String("return-values", "", "Use `ReturnValues` if you want to get the item attributes as they appeared before they were deleted.")
		dynamodb_deleteItemCmd.Flags().String("return-values-on-condition-check-failure", "", "An optional parameter that returns the item attributes for a `DeleteItem` operation that failed a condition check.")
		dynamodb_deleteItemCmd.Flags().String("table-name", "", "The name of the table from which to delete the item.")
		dynamodb_deleteItemCmd.MarkFlagRequired("key")
		dynamodb_deleteItemCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_deleteItemCmd)
}
