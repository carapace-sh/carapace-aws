package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_executeStatementCmd = &cobra.Command{
	Use:   "execute-statement",
	Short: "Runs a SQL statement against a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_executeStatementCmd).Standalone()

	rdsData_executeStatementCmd.Flags().Bool("continue-after-timeout", false, "A value that indicates whether to continue running the statement after the call times out.")
	rdsData_executeStatementCmd.Flags().String("database", "", "The name of the database.")
	rdsData_executeStatementCmd.Flags().String("format-records-as", "", "A value that indicates whether to format the result set as a single JSON string.")
	rdsData_executeStatementCmd.Flags().Bool("include-result-metadata", false, "A value that indicates whether to include metadata in the results.")
	rdsData_executeStatementCmd.Flags().Bool("no-continue-after-timeout", false, "A value that indicates whether to continue running the statement after the call times out.")
	rdsData_executeStatementCmd.Flags().Bool("no-include-result-metadata", false, "A value that indicates whether to include metadata in the results.")
	rdsData_executeStatementCmd.Flags().String("parameters", "", "The parameters for the SQL statement.")
	rdsData_executeStatementCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.")
	rdsData_executeStatementCmd.Flags().String("result-set-options", "", "Options that control how the result set is returned.")
	rdsData_executeStatementCmd.Flags().String("schema", "", "The name of the database schema.")
	rdsData_executeStatementCmd.Flags().String("secret-arn", "", "The ARN of the secret that enables access to the DB cluster.")
	rdsData_executeStatementCmd.Flags().String("sql", "", "The SQL statement to run.")
	rdsData_executeStatementCmd.Flags().String("transaction-id", "", "The identifier of a transaction that was started by using the `BeginTransaction` operation.")
	rdsData_executeStatementCmd.Flag("no-continue-after-timeout").Hidden = true
	rdsData_executeStatementCmd.Flag("no-include-result-metadata").Hidden = true
	rdsData_executeStatementCmd.MarkFlagRequired("resource-arn")
	rdsData_executeStatementCmd.MarkFlagRequired("secret-arn")
	rdsData_executeStatementCmd.MarkFlagRequired("sql")
	rdsDataCmd.AddCommand(rdsData_executeStatementCmd)
}
