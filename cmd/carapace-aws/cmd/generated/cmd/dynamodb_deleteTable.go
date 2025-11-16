package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "The `DeleteTable` operation deletes a table and all of its items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_deleteTableCmd).Standalone()

	dynamodb_deleteTableCmd.Flags().String("table-name", "", "The name of the table to delete.")
	dynamodb_deleteTableCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_deleteTableCmd)
}
