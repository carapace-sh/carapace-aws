package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_getItemCmd = &cobra.Command{
	Use:   "get-item",
	Short: "The `GetItem` operation returns a set of attributes for the item with the given primary key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_getItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_getItemCmd).Standalone()

		dynamodb_getItemCmd.Flags().String("attributes-to-get", "", "This is a legacy parameter.")
		dynamodb_getItemCmd.Flags().String("consistent-read", "", "Determines the read consistency model: If set to `true`, then the operation uses strongly consistent reads; otherwise, the operation uses eventually consistent reads.")
		dynamodb_getItemCmd.Flags().String("expression-attribute-names", "", "One or more substitution tokens for attribute names in an expression.")
		dynamodb_getItemCmd.Flags().String("key", "", "A map of attribute names to `AttributeValue` objects, representing the primary key of the item to retrieve.")
		dynamodb_getItemCmd.Flags().String("projection-expression", "", "A string that identifies one or more attributes to retrieve from the table.")
		dynamodb_getItemCmd.Flags().String("return-consumed-capacity", "", "")
		dynamodb_getItemCmd.Flags().String("table-name", "", "The name of the table containing the requested item.")
		dynamodb_getItemCmd.MarkFlagRequired("key")
		dynamodb_getItemCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_getItemCmd)
}
