package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_batchGetQueryExecutionCmd = &cobra.Command{
	Use:   "batch-get-query-execution",
	Short: "Returns the details of a single query execution or a list of up to 50 query executions, which you provide as an array of query execution ID strings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_batchGetQueryExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_batchGetQueryExecutionCmd).Standalone()

		athena_batchGetQueryExecutionCmd.Flags().String("query-execution-ids", "", "An array of query execution IDs.")
		athena_batchGetQueryExecutionCmd.MarkFlagRequired("query-execution-ids")
	})
	athenaCmd.AddCommand(athena_batchGetQueryExecutionCmd)
}
