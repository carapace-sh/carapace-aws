package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateItemCmd = &cobra.Command{
	Use:   "update-item",
	Short: "Edits an existing item's attributes, or adds a new item to the table if it does not already exist.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateItemCmd).Standalone()

		dynamodb_updateItemCmd.Flags().String("attribute-updates", "", "This is a legacy parameter.")
		dynamodb_updateItemCmd.Flags().String("condition-expression", "", "A condition that must be satisfied in order for a conditional update to succeed.")
		dynamodb_updateItemCmd.Flags().String("conditional-operator", "", "This is a legacy parameter.")
		dynamodb_updateItemCmd.Flags().String("expected", "", "This is a legacy parameter.")
		dynamodb_updateItemCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_updateItemCmd.Flags().String("expression-attribute-values", "", "One or more values that can be substituted in an expression.")
		dynamodb_updateItemCmd.Flags().String("key", "", "The primary key of the item to be updated.")
		dynamodb_updateItemCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_updateItemCmd.Flags().String("return-item-collection-metrics", "", "Determines whether item collection metrics are returned.")
		dynamodb_updateItemCmd.Flags().String("return-values", "", "Use `ReturnValues` if you want to get the item attributes as they appear before or after they are successfully updated.")
		dynamodb_updateItemCmd.Flags().String("return-values-on-condition-check-failure", "", "An optional parameter that returns the item attributes for an `UpdateItem` operation that failed a condition check.")
		dynamodb_updateItemCmd.Flags().String("table-name", "", "The name of the table containing the item to update.")
		dynamodb_updateItemCmd.Flags().String("update-expression", "", "An expression that defines one or more attributes to be updated, the action to be performed on them, and new values for them.")
		dynamodb_updateItemCmd.MarkFlagRequired("key")
		dynamodb_updateItemCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_updateItemCmd)
}
