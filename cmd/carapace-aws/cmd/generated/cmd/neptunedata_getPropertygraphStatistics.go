package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getPropertygraphStatisticsCmd = &cobra.Command{
	Use:   "get-propertygraph-statistics",
	Short: "Gets property graph statistics (Gremlin and openCypher).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getPropertygraphStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getPropertygraphStatisticsCmd).Standalone()

	})
	neptunedataCmd.AddCommand(neptunedata_getPropertygraphStatisticsCmd)
}
