package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_commitTransactionCmd = &cobra.Command{
	Use:   "commit-transaction",
	Short: "Ends a SQL transaction started with the `BeginTransaction` operation and commits the changes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_commitTransactionCmd).Standalone()

	rdsData_commitTransactionCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.")
	rdsData_commitTransactionCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the DB cluster.")
	rdsData_commitTransactionCmd.Flags().String("transaction-id", "", "The identifier of the transaction to end and commit.")
	rdsData_commitTransactionCmd.MarkFlagRequired("resource-arn")
	rdsData_commitTransactionCmd.MarkFlagRequired("secret-arn")
	rdsData_commitTransactionCmd.MarkFlagRequired("transaction-id")
	rdsDataCmd.AddCommand(rdsData_commitTransactionCmd)
}
