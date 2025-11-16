package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getOpenCypherQueryStatusCmd = &cobra.Command{
	Use:   "get-open-cypher-query-status",
	Short: "Retrieves the status of a specified openCypher query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getOpenCypherQueryStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getOpenCypherQueryStatusCmd).Standalone()

		neptunedata_getOpenCypherQueryStatusCmd.Flags().String("query-id", "", "The unique ID of the openCypher query for which to retrieve the query status.")
		neptunedata_getOpenCypherQueryStatusCmd.MarkFlagRequired("query-id")
	})
	neptunedataCmd.AddCommand(neptunedata_getOpenCypherQueryStatusCmd)
}
