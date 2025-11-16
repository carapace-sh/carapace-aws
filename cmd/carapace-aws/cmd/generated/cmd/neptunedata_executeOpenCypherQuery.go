package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeOpenCypherQueryCmd = &cobra.Command{
	Use:   "execute-open-cypher-query",
	Short: "Executes an openCypher query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeOpenCypherQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_executeOpenCypherQueryCmd).Standalone()

		neptunedata_executeOpenCypherQueryCmd.Flags().String("open-cypher-query", "", "The openCypher query string to be executed.")
		neptunedata_executeOpenCypherQueryCmd.Flags().String("parameters", "", "The openCypher query parameters for query execution.")
		neptunedata_executeOpenCypherQueryCmd.MarkFlagRequired("open-cypher-query")
	})
	neptunedataCmd.AddCommand(neptunedata_executeOpenCypherQueryCmd)
}
