package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getPreparedStatementCmd = &cobra.Command{
	Use:   "get-prepared-statement",
	Short: "Retrieves the prepared statement with the specified name from the specified workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getPreparedStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getPreparedStatementCmd).Standalone()

		athena_getPreparedStatementCmd.Flags().String("statement-name", "", "The name of the prepared statement to retrieve.")
		athena_getPreparedStatementCmd.Flags().String("work-group", "", "The workgroup to which the statement to be retrieved belongs.")
		athena_getPreparedStatementCmd.MarkFlagRequired("statement-name")
		athena_getPreparedStatementCmd.MarkFlagRequired("work-group")
	})
	athenaCmd.AddCommand(athena_getPreparedStatementCmd)
}
