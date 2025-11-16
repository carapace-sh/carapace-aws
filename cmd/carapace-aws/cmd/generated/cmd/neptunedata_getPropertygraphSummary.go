package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getPropertygraphSummaryCmd = &cobra.Command{
	Use:   "get-propertygraph-summary",
	Short: "Gets a graph summary for a property graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getPropertygraphSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getPropertygraphSummaryCmd).Standalone()

		neptunedata_getPropertygraphSummaryCmd.Flags().String("mode", "", "Mode can take one of two values: `BASIC` (the default), and `DETAILED`.")
	})
	neptunedataCmd.AddCommand(neptunedata_getPropertygraphSummaryCmd)
}
