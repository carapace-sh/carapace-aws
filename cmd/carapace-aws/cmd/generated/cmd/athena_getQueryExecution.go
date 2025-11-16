package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getQueryExecutionCmd = &cobra.Command{
	Use:   "get-query-execution",
	Short: "Returns information about a single execution of a query if you have access to the workgroup in which the query ran.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getQueryExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getQueryExecutionCmd).Standalone()

		athena_getQueryExecutionCmd.Flags().String("query-execution-id", "", "The unique ID of the query execution.")
		athena_getQueryExecutionCmd.MarkFlagRequired("query-execution-id")
	})
	athenaCmd.AddCommand(athena_getQueryExecutionCmd)
}
