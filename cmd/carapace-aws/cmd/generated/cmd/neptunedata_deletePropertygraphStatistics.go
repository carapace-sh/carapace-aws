package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_deletePropertygraphStatisticsCmd = &cobra.Command{
	Use:   "delete-propertygraph-statistics",
	Short: "Deletes statistics for Gremlin and openCypher (property graph) data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_deletePropertygraphStatisticsCmd).Standalone()

	neptunedataCmd.AddCommand(neptunedata_deletePropertygraphStatisticsCmd)
}
