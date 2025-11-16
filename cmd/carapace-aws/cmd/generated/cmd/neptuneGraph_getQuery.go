package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getQueryCmd = &cobra.Command{
	Use:   "get-query",
	Short: "Retrieves the status of a specified query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_getQueryCmd).Standalone()

		neptuneGraph_getQueryCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_getQueryCmd.Flags().String("query-id", "", "The ID of the query in question.")
		neptuneGraph_getQueryCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_getQueryCmd.MarkFlagRequired("query-id")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_getQueryCmd)
}
