package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_getStatementResultCmd = &cobra.Command{
	Use:   "get-statement-result",
	Short: "Fetches the temporarily cached result of an SQL statement in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_getStatementResultCmd).Standalone()

	redshiftData_getStatementResultCmd.Flags().String("id", "", "The identifier of the SQL statement whose results are to be fetched.")
	redshiftData_getStatementResultCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
	redshiftData_getStatementResultCmd.MarkFlagRequired("id")
	redshiftDataCmd.AddCommand(redshiftData_getStatementResultCmd)
}
