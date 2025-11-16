package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_batchExecuteStatementCmd = &cobra.Command{
	Use:   "batch-execute-statement",
	Short: "This operation allows you to perform batch reads or writes on data stored in DynamoDB, using PartiQL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_batchExecuteStatementCmd).Standalone()

	dynamodb_batchExecuteStatementCmd.Flags().String("return-consumed-capacity", "", "")
	dynamodb_batchExecuteStatementCmd.Flags().String("statements", "", "The list of PartiQL statements representing the batch to run.")
	dynamodb_batchExecuteStatementCmd.MarkFlagRequired("statements")
	dynamodbCmd.AddCommand(dynamodb_batchExecuteStatementCmd)
}
