package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_executeQueryCmd = &cobra.Command{
	Use:   "execute-query",
	Short: "Execute an openCypher query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_executeQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_executeQueryCmd).Standalone()

		neptuneGraph_executeQueryCmd.Flags().String("explain-mode", "", "The explain mode parameter returns a query explain instead of the actual query results.")
		neptuneGraph_executeQueryCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_executeQueryCmd.Flags().String("language", "", "The query language the query is written in.")
		neptuneGraph_executeQueryCmd.Flags().String("parameters", "", "The data parameters the query can use in JSON format.")
		neptuneGraph_executeQueryCmd.Flags().String("plan-cache", "", "Query plan cache is a feature that saves the query plan and reuses it on successive executions of the same query.")
		neptuneGraph_executeQueryCmd.Flags().String("query-string", "", "The query string to be executed.")
		neptuneGraph_executeQueryCmd.Flags().String("query-timeout-milliseconds", "", "Specifies the query timeout duration, in milliseconds.")
		neptuneGraph_executeQueryCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_executeQueryCmd.MarkFlagRequired("language")
		neptuneGraph_executeQueryCmd.MarkFlagRequired("query-string")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_executeQueryCmd)
}
