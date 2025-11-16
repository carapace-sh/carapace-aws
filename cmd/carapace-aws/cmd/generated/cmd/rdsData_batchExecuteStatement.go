package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_batchExecuteStatementCmd = &cobra.Command{
	Use:   "batch-execute-statement",
	Short: "Runs a batch SQL statement over an array of data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_batchExecuteStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsData_batchExecuteStatementCmd).Standalone()

		rdsData_batchExecuteStatementCmd.Flags().String("database", "", "The name of the database.")
		rdsData_batchExecuteStatementCmd.Flags().String("parameter-sets", "", "The parameter set for the batch operation.")
		rdsData_batchExecuteStatementCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.")
		rdsData_batchExecuteStatementCmd.Flags().String("schema", "", "The name of the database schema.")
		rdsData_batchExecuteStatementCmd.Flags().String("secret-arn", "", "The ARN of the secret that enables access to the DB cluster.")
		rdsData_batchExecuteStatementCmd.Flags().String("sql", "", "The SQL statement to run.")
		rdsData_batchExecuteStatementCmd.Flags().String("transaction-id", "", "The identifier of a transaction that was started by using the `BeginTransaction` operation.")
		rdsData_batchExecuteStatementCmd.MarkFlagRequired("resource-arn")
		rdsData_batchExecuteStatementCmd.MarkFlagRequired("secret-arn")
		rdsData_batchExecuteStatementCmd.MarkFlagRequired("sql")
	})
	rdsDataCmd.AddCommand(rdsData_batchExecuteStatementCmd)
}
