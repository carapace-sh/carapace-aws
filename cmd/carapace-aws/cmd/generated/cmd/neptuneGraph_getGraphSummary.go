package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getGraphSummaryCmd = &cobra.Command{
	Use:   "get-graph-summary",
	Short: "Gets a graph summary for a property graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getGraphSummaryCmd).Standalone()

	neptuneGraph_getGraphSummaryCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_getGraphSummaryCmd.Flags().String("mode", "", "The summary mode can take one of two values: `basic` (the default), and `detailed`.")
	neptuneGraph_getGraphSummaryCmd.MarkFlagRequired("graph-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_getGraphSummaryCmd)
}
