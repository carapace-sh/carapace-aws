package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "Returns an array of table names associated with the current account and endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listTablesCmd).Standalone()

	dynamodb_listTablesCmd.Flags().String("exclusive-start-table-name", "", "The first table name that this operation will evaluate.")
	dynamodb_listTablesCmd.Flags().String("limit", "", "A maximum number of table names to return.")
	dynamodbCmd.AddCommand(dynamodb_listTablesCmd)
}
