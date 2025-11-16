package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getRdfgraphSummaryCmd = &cobra.Command{
	Use:   "get-rdfgraph-summary",
	Short: "Gets a graph summary for an RDF graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getRdfgraphSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getRdfgraphSummaryCmd).Standalone()

		neptunedata_getRdfgraphSummaryCmd.Flags().String("mode", "", "Mode can take one of two values: `BASIC` (the default), and `DETAILED`.")
	})
	neptunedataCmd.AddCommand(neptunedata_getRdfgraphSummaryCmd)
}
