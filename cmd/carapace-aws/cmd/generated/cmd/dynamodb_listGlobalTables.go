package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listGlobalTablesCmd = &cobra.Command{
	Use:   "list-global-tables",
	Short: "Lists all global tables that have a replica in the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listGlobalTablesCmd).Standalone()

	dynamodb_listGlobalTablesCmd.Flags().String("exclusive-start-global-table-name", "", "The first global table name that this operation will evaluate.")
	dynamodb_listGlobalTablesCmd.Flags().String("limit", "", "The maximum number of table names to return, if the parameter is not specified DynamoDB defaults to 100.")
	dynamodb_listGlobalTablesCmd.Flags().String("region-name", "", "Lists the global tables in a specific Region.")
	dynamodbCmd.AddCommand(dynamodb_listGlobalTablesCmd)
}
