package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_executeTransactionCmd = &cobra.Command{
	Use:   "execute-transaction",
	Short: "This operation allows you to perform transactional reads or writes on data stored in DynamoDB, using PartiQL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_executeTransactionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_executeTransactionCmd).Standalone()

		dynamodb_executeTransactionCmd.Flags().String("client-request-token", "", "Set this value to get remaining results, if `NextToken` was returned in the statement response.")
		dynamodb_executeTransactionCmd.Flags().String("return-consumed-capacity", "", "Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response.")
		dynamodb_executeTransactionCmd.Flags().String("transact-statements", "", "The list of PartiQL statements representing the transaction to run.")
		dynamodb_executeTransactionCmd.MarkFlagRequired("transact-statements")
	})
	dynamodbCmd.AddCommand(dynamodb_executeTransactionCmd)
}
