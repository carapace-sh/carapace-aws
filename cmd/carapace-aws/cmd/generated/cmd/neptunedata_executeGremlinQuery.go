package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeGremlinQueryCmd = &cobra.Command{
	Use:   "execute-gremlin-query",
	Short: "This commands executes a Gremlin query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeGremlinQueryCmd).Standalone()

	neptunedata_executeGremlinQueryCmd.Flags().String("gremlin-query", "", "Using this API, you can run Gremlin queries in string format much as you can using the HTTP endpoint.")
	neptunedata_executeGremlinQueryCmd.Flags().String("serializer", "", "If non-null, the query results are returned in a serialized response message in the format specified by this parameter.")
	neptunedata_executeGremlinQueryCmd.MarkFlagRequired("gremlin-query")
	neptunedataCmd.AddCommand(neptunedata_executeGremlinQueryCmd)
}
