package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listOpenCypherQueriesCmd = &cobra.Command{
	Use:   "list-open-cypher-queries",
	Short: "Lists active openCypher queries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listOpenCypherQueriesCmd).Standalone()

	neptunedata_listOpenCypherQueriesCmd.Flags().Bool("include-waiting", false, "When set to `TRUE` and other parameters are not present, causes status information to be returned for waiting queries as well as for running queries.")
	neptunedata_listOpenCypherQueriesCmd.Flags().Bool("no-include-waiting", false, "When set to `TRUE` and other parameters are not present, causes status information to be returned for waiting queries as well as for running queries.")
	neptunedata_listOpenCypherQueriesCmd.Flag("no-include-waiting").Hidden = true
	neptunedataCmd.AddCommand(neptunedata_listOpenCypherQueriesCmd)
}
