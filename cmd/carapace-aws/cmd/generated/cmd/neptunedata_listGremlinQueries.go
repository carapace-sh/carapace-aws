package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listGremlinQueriesCmd = &cobra.Command{
	Use:   "list-gremlin-queries",
	Short: "Lists active Gremlin queries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listGremlinQueriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_listGremlinQueriesCmd).Standalone()

		neptunedata_listGremlinQueriesCmd.Flags().Bool("include-waiting", false, "If set to `TRUE`, the list returned includes waiting queries.")
		neptunedata_listGremlinQueriesCmd.Flags().Bool("no-include-waiting", false, "If set to `TRUE`, the list returned includes waiting queries.")
		neptunedata_listGremlinQueriesCmd.Flag("no-include-waiting").Hidden = true
	})
	neptunedataCmd.AddCommand(neptunedata_listGremlinQueriesCmd)
}
