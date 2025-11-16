package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listQueriesCmd = &cobra.Command{
	Use:   "list-queries",
	Short: "Lists active openCypher queries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listQueriesCmd).Standalone()

	neptuneGraph_listQueriesCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_listQueriesCmd.Flags().String("max-results", "", "The maximum number of results to be fetched by the API.")
	neptuneGraph_listQueriesCmd.Flags().String("state", "", "Filtered list of queries based on state.")
	neptuneGraph_listQueriesCmd.MarkFlagRequired("graph-identifier")
	neptuneGraph_listQueriesCmd.MarkFlagRequired("max-results")
	neptuneGraphCmd.AddCommand(neptuneGraph_listQueriesCmd)
}
