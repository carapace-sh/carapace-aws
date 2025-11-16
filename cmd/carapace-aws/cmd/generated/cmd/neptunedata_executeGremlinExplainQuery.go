package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeGremlinExplainQueryCmd = &cobra.Command{
	Use:   "execute-gremlin-explain-query",
	Short: "Executes a Gremlin Explain query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeGremlinExplainQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_executeGremlinExplainQueryCmd).Standalone()

		neptunedata_executeGremlinExplainQueryCmd.Flags().String("gremlin-query", "", "The Gremlin explain query string.")
		neptunedata_executeGremlinExplainQueryCmd.MarkFlagRequired("gremlin-query")
	})
	neptunedataCmd.AddCommand(neptunedata_executeGremlinExplainQueryCmd)
}
