package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_executeStatementCmd = &cobra.Command{
	Use:   "execute-statement",
	Short: "Runs an SQL statement, which can be data manipulation language (DML) or data definition language (DDL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_executeStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_executeStatementCmd).Standalone()

		redshiftData_executeStatementCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		redshiftData_executeStatementCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshiftData_executeStatementCmd.Flags().String("database", "", "The name of the database.")
		redshiftData_executeStatementCmd.Flags().String("db-user", "", "The database user name.")
		redshiftData_executeStatementCmd.Flags().Bool("no-with-event", false, "A value that indicates whether to send an event to the Amazon EventBridge event bus after the SQL statement runs.")
		redshiftData_executeStatementCmd.Flags().String("parameters", "", "The parameters for the SQL statement.")
		redshiftData_executeStatementCmd.Flags().String("result-format", "", "The data format of the result of the SQL statement.")
		redshiftData_executeStatementCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
		redshiftData_executeStatementCmd.Flags().String("session-id", "", "The session identifier of the query.")
		redshiftData_executeStatementCmd.Flags().String("session-keep-alive-seconds", "", "The number of seconds to keep the session alive after the query finishes.")
		redshiftData_executeStatementCmd.Flags().String("sql", "", "The SQL statement text to run.")
		redshiftData_executeStatementCmd.Flags().String("statement-name", "", "The name of the SQL statement.")
		redshiftData_executeStatementCmd.Flags().Bool("with-event", false, "A value that indicates whether to send an event to the Amazon EventBridge event bus after the SQL statement runs.")
		redshiftData_executeStatementCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
		redshiftData_executeStatementCmd.Flag("no-with-event").Hidden = true
		redshiftData_executeStatementCmd.MarkFlagRequired("sql")
	})
	redshiftDataCmd.AddCommand(redshiftData_executeStatementCmd)
}
