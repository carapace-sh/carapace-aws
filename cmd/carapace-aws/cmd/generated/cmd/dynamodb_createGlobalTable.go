package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_createGlobalTableCmd = &cobra.Command{
	Use:   "create-global-table",
	Short: "Creates a global table from an existing table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_createGlobalTableCmd).Standalone()

	dynamodb_createGlobalTableCmd.Flags().String("global-table-name", "", "The global table name.")
	dynamodb_createGlobalTableCmd.Flags().String("replication-group", "", "The Regions where the global table needs to be created.")
	dynamodb_createGlobalTableCmd.MarkFlagRequired("global-table-name")
	dynamodb_createGlobalTableCmd.MarkFlagRequired("replication-group")
	dynamodbCmd.AddCommand(dynamodb_createGlobalTableCmd)
}
