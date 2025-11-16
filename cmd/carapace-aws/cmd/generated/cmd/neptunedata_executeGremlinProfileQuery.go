package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeGremlinProfileQueryCmd = &cobra.Command{
	Use:   "execute-gremlin-profile-query",
	Short: "Executes a Gremlin Profile query, which runs a specified traversal, collects various metrics about the run, and produces a profile report as output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeGremlinProfileQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_executeGremlinProfileQueryCmd).Standalone()

		neptunedata_executeGremlinProfileQueryCmd.Flags().String("chop", "", "If non-zero, causes the results string to be truncated at that number of characters.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().String("gremlin-query", "", "The Gremlin query string to profile.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().Bool("index-ops", false, "If this flag is set to `TRUE`, the results include a detailed report of all index operations that took place during query execution and serialization.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().Bool("no-index-ops", false, "If this flag is set to `TRUE`, the results include a detailed report of all index operations that took place during query execution and serialization.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().Bool("no-results", false, "If this flag is set to `TRUE`, the query results are gathered and displayed as part of the profile report.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().Bool("results", false, "If this flag is set to `TRUE`, the query results are gathered and displayed as part of the profile report.")
		neptunedata_executeGremlinProfileQueryCmd.Flags().String("serializer", "", "If non-null, the gathered results are returned in a serialized response message in the format specified by this parameter.")
		neptunedata_executeGremlinProfileQueryCmd.MarkFlagRequired("gremlin-query")
		neptunedata_executeGremlinProfileQueryCmd.Flag("no-index-ops").Hidden = true
		neptunedata_executeGremlinProfileQueryCmd.Flag("no-results").Hidden = true
	})
	neptunedataCmd.AddCommand(neptunedata_executeGremlinProfileQueryCmd)
}
