package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createPreparedStatementCmd = &cobra.Command{
	Use:   "create-prepared-statement",
	Short: "Creates a prepared statement for use with SQL queries in Athena.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createPreparedStatementCmd).Standalone()

	athena_createPreparedStatementCmd.Flags().String("description", "", "The description of the prepared statement.")
	athena_createPreparedStatementCmd.Flags().String("query-statement", "", "The query string for the prepared statement.")
	athena_createPreparedStatementCmd.Flags().String("statement-name", "", "The name of the prepared statement.")
	athena_createPreparedStatementCmd.Flags().String("work-group", "", "The name of the workgroup to which the prepared statement belongs.")
	athena_createPreparedStatementCmd.MarkFlagRequired("query-statement")
	athena_createPreparedStatementCmd.MarkFlagRequired("statement-name")
	athena_createPreparedStatementCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_createPreparedStatementCmd)
}
