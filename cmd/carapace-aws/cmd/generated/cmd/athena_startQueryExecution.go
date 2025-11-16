package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_startQueryExecutionCmd = &cobra.Command{
	Use:   "start-query-execution",
	Short: "Runs the SQL query statements contained in the `Query`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_startQueryExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_startQueryExecutionCmd).Standalone()

		athena_startQueryExecutionCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the query is idempotent (executes only once).")
		athena_startQueryExecutionCmd.Flags().String("execution-parameters", "", "A list of values for the parameters in a query.")
		athena_startQueryExecutionCmd.Flags().String("query-execution-context", "", "The database within which the query executes.")
		athena_startQueryExecutionCmd.Flags().String("query-string", "", "The SQL query statements to be executed.")
		athena_startQueryExecutionCmd.Flags().String("result-configuration", "", "Specifies information about where and how to save the results of the query execution.")
		athena_startQueryExecutionCmd.Flags().String("result-reuse-configuration", "", "Specifies the query result reuse behavior for the query.")
		athena_startQueryExecutionCmd.Flags().String("work-group", "", "The name of the workgroup in which the query is being started.")
		athena_startQueryExecutionCmd.MarkFlagRequired("query-string")
	})
	athenaCmd.AddCommand(athena_startQueryExecutionCmd)
}
