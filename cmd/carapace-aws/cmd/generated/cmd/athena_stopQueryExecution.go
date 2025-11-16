package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_stopQueryExecutionCmd = &cobra.Command{
	Use:   "stop-query-execution",
	Short: "Stops a query execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_stopQueryExecutionCmd).Standalone()

	athena_stopQueryExecutionCmd.Flags().String("query-execution-id", "", "The unique ID of the query execution to stop.")
	athena_stopQueryExecutionCmd.MarkFlagRequired("query-execution-id")
	athenaCmd.AddCommand(athena_stopQueryExecutionCmd)
}
