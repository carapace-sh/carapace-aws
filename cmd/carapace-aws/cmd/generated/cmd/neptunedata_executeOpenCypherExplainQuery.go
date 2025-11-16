package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeOpenCypherExplainQueryCmd = &cobra.Command{
	Use:   "execute-open-cypher-explain-query",
	Short: "Executes an openCypher `explain` request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeOpenCypherExplainQueryCmd).Standalone()

	neptunedata_executeOpenCypherExplainQueryCmd.Flags().String("explain-mode", "", "The openCypher `explain` mode.")
	neptunedata_executeOpenCypherExplainQueryCmd.Flags().String("open-cypher-query", "", "The openCypher query string.")
	neptunedata_executeOpenCypherExplainQueryCmd.Flags().String("parameters", "", "The openCypher query parameters.")
	neptunedata_executeOpenCypherExplainQueryCmd.MarkFlagRequired("explain-mode")
	neptunedata_executeOpenCypherExplainQueryCmd.MarkFlagRequired("open-cypher-query")
	neptunedataCmd.AddCommand(neptunedata_executeOpenCypherExplainQueryCmd)
}
