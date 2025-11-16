package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deletePreparedStatementCmd = &cobra.Command{
	Use:   "delete-prepared-statement",
	Short: "Deletes the prepared statement with the specified name from the specified workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deletePreparedStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_deletePreparedStatementCmd).Standalone()

		athena_deletePreparedStatementCmd.Flags().String("statement-name", "", "The name of the prepared statement to delete.")
		athena_deletePreparedStatementCmd.Flags().String("work-group", "", "The workgroup to which the statement to be deleted belongs.")
		athena_deletePreparedStatementCmd.MarkFlagRequired("statement-name")
		athena_deletePreparedStatementCmd.MarkFlagRequired("work-group")
	})
	athenaCmd.AddCommand(athena_deletePreparedStatementCmd)
}
