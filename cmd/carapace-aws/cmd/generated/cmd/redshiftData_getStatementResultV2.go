package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_getStatementResultV2Cmd = &cobra.Command{
	Use:   "get-statement-result-v2",
	Short: "Fetches the temporarily cached result of an SQL statement in CSV format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_getStatementResultV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_getStatementResultV2Cmd).Standalone()

		redshiftData_getStatementResultV2Cmd.Flags().String("id", "", "The identifier of the SQL statement whose results are to be fetched.")
		redshiftData_getStatementResultV2Cmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshiftData_getStatementResultV2Cmd.MarkFlagRequired("id")
	})
	redshiftDataCmd.AddCommand(redshiftData_getStatementResultV2Cmd)
}
