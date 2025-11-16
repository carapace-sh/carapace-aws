package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_executeStatementCmd = &cobra.Command{
	Use:   "execute-statement",
	Short: "This operation allows you to perform reads and singleton writes on data stored in DynamoDB, using PartiQL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_executeStatementCmd).Standalone()

	dynamodb_executeStatementCmd.Flags().String("consistent-read", "", "The consistency of a read operation.")
	dynamodb_executeStatementCmd.Flags().String("limit", "", "The maximum number of items to evaluate (not necessarily the number of matching items).")
	dynamodb_executeStatementCmd.Flags().String("next-token", "", "Set this value to get remaining results, if `NextToken` was returned in the statement response.")
	dynamodb_executeStatementCmd.Flags().String("parameters", "", "The parameters for the PartiQL statement, if any.")
	dynamodb_executeStatementCmd.Flags().String("return-consumed-capacity", "", "")
	dynamodb_executeStatementCmd.Flags().String("return-values-on-condition-check-failure", "", "An optional parameter that returns the item attributes for an `ExecuteStatement` operation that failed a condition check.")
	dynamodb_executeStatementCmd.Flags().String("statement", "", "The PartiQL statement representing the operation to run.")
	dynamodb_executeStatementCmd.MarkFlagRequired("statement")
	dynamodbCmd.AddCommand(dynamodb_executeStatementCmd)
}
