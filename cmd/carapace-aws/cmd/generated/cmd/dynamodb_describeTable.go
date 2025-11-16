package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeTableCmd = &cobra.Command{
	Use:   "describe-table",
	Short: "Returns information about the table, including the current status of the table, when it was created, the primary key schema, and any indexes on the table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeTableCmd).Standalone()

	dynamodb_describeTableCmd.Flags().String("table-name", "", "The name of the table to describe.")
	dynamodb_describeTableCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_describeTableCmd)
}
