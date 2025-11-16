package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_rollbackTransactionCmd = &cobra.Command{
	Use:   "rollback-transaction",
	Short: "Performs a rollback of a transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_rollbackTransactionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsData_rollbackTransactionCmd).Standalone()

		rdsData_rollbackTransactionCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.")
		rdsData_rollbackTransactionCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the DB cluster.")
		rdsData_rollbackTransactionCmd.Flags().String("transaction-id", "", "The identifier of the transaction to roll back.")
		rdsData_rollbackTransactionCmd.MarkFlagRequired("resource-arn")
		rdsData_rollbackTransactionCmd.MarkFlagRequired("secret-arn")
		rdsData_rollbackTransactionCmd.MarkFlagRequired("transaction-id")
	})
	rdsDataCmd.AddCommand(rdsData_rollbackTransactionCmd)
}
