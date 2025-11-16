package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_batchExecuteStatementCmd = &cobra.Command{
	Use:   "batch-execute-statement",
	Short: "Runs one or more SQL statements, which can be data manipulation language (DML) or data definition language (DDL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_batchExecuteStatementCmd).Standalone()

	redshiftData_batchExecuteStatementCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	redshiftData_batchExecuteStatementCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
	redshiftData_batchExecuteStatementCmd.Flags().String("database", "", "The name of the database.")
	redshiftData_batchExecuteStatementCmd.Flags().String("db-user", "", "The database user name.")
	redshiftData_batchExecuteStatementCmd.Flags().Bool("no-with-event", false, "A value that indicates whether to send an event to the Amazon EventBridge event bus after the SQL statements run.")
	redshiftData_batchExecuteStatementCmd.Flags().String("result-format", "", "The data format of the result of the SQL statement.")
	redshiftData_batchExecuteStatementCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
	redshiftData_batchExecuteStatementCmd.Flags().String("session-id", "", "The session identifier of the query.")
	redshiftData_batchExecuteStatementCmd.Flags().String("session-keep-alive-seconds", "", "The number of seconds to keep the session alive after the query finishes.")
	redshiftData_batchExecuteStatementCmd.Flags().String("sqls", "", "One or more SQL statements to run.")
	redshiftData_batchExecuteStatementCmd.Flags().String("statement-name", "", "The name of the SQL statements.")
	redshiftData_batchExecuteStatementCmd.Flags().Bool("with-event", false, "A value that indicates whether to send an event to the Amazon EventBridge event bus after the SQL statements run.")
	redshiftData_batchExecuteStatementCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
	redshiftData_batchExecuteStatementCmd.Flag("no-with-event").Hidden = true
	redshiftData_batchExecuteStatementCmd.MarkFlagRequired("sqls")
	redshiftDataCmd.AddCommand(redshiftData_batchExecuteStatementCmd)
}
