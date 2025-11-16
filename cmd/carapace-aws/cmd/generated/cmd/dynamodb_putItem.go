package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_putItemCmd = &cobra.Command{
	Use:   "put-item",
	Short: "Creates a new item, or replaces an old item with a new item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_putItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_putItemCmd).Standalone()

		dynamodb_putItemCmd.Flags().String("condition-expression", "", "A condition that must be satisfied in order for a conditional `PutItem` operation to succeed.")
		dynamodb_putItemCmd.Flags().String("conditional-operator", "", "This is a legacy parameter.")
		dynamodb_putItemCmd.Flags().String("expected", "", "This is a legacy parameter.")
		dynamodb_putItemCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_putItemCmd.Flags().String("expression-attribute-values", "", "One or more values that can be substituted in an expression.")
		dynamodb_putItemCmd.Flags().String("item", "", "A map of attribute name/value pairs, one for each attribute.")
		dynamodb_putItemCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_putItemCmd.Flags().String("return-item-collection-metrics", "", "Determines whether item collection metrics are returned.")
		dynamodb_putItemCmd.Flags().String("return-values", "", "Use `ReturnValues` if you want to get the item attributes as they appeared before they were updated with the `PutItem` request.")
		dynamodb_putItemCmd.Flags().String("return-values-on-condition-check-failure", "", "An optional parameter that returns the item attributes for a `PutItem` operation that failed a condition check.")
		dynamodb_putItemCmd.Flags().String("table-name", "", "The name of the table to contain the item.")
		dynamodb_putItemCmd.MarkFlagRequired("item")
		dynamodb_putItemCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_putItemCmd)
}
