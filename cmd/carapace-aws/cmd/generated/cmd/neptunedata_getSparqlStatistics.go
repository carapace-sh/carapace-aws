package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getSparqlStatisticsCmd = &cobra.Command{
	Use:   "get-sparql-statistics",
	Short: "Gets RDF statistics (SPARQL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getSparqlStatisticsCmd).Standalone()

	neptunedataCmd.AddCommand(neptunedata_getSparqlStatisticsCmd)
}
