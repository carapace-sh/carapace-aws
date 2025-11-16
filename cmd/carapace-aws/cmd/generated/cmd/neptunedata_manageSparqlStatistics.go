package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_manageSparqlStatisticsCmd = &cobra.Command{
	Use:   "manage-sparql-statistics",
	Short: "Manages the generation and use of RDF graph statistics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_manageSparqlStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_manageSparqlStatisticsCmd).Standalone()

		neptunedata_manageSparqlStatisticsCmd.Flags().String("mode", "", "The statistics generation mode.")
	})
	neptunedataCmd.AddCommand(neptunedata_manageSparqlStatisticsCmd)
}
