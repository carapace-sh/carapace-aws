package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_managePropertygraphStatisticsCmd = &cobra.Command{
	Use:   "manage-propertygraph-statistics",
	Short: "Manages the generation and use of property graph statistics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_managePropertygraphStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_managePropertygraphStatisticsCmd).Standalone()

		neptunedata_managePropertygraphStatisticsCmd.Flags().String("mode", "", "The statistics generation mode.")
	})
	neptunedataCmd.AddCommand(neptunedata_managePropertygraphStatisticsCmd)
}
