package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updatePreparedStatementCmd = &cobra.Command{
	Use:   "update-prepared-statement",
	Short: "Updates a prepared statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updatePreparedStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_updatePreparedStatementCmd).Standalone()

		athena_updatePreparedStatementCmd.Flags().String("description", "", "The description of the prepared statement.")
		athena_updatePreparedStatementCmd.Flags().String("query-statement", "", "The query string for the prepared statement.")
		athena_updatePreparedStatementCmd.Flags().String("statement-name", "", "The name of the prepared statement.")
		athena_updatePreparedStatementCmd.Flags().String("work-group", "", "The workgroup for the prepared statement.")
		athena_updatePreparedStatementCmd.MarkFlagRequired("query-statement")
		athena_updatePreparedStatementCmd.MarkFlagRequired("statement-name")
		athena_updatePreparedStatementCmd.MarkFlagRequired("work-group")
	})
	athenaCmd.AddCommand(athena_updatePreparedStatementCmd)
}
