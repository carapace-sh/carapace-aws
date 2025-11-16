package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listSignalCatalogNodesCmd = &cobra.Command{
	Use:   "list-signal-catalog-nodes",
	Short: "Lists of information about the signals (nodes) specified in a signal catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listSignalCatalogNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listSignalCatalogNodesCmd).Standalone()

		iotfleetwise_listSignalCatalogNodesCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listSignalCatalogNodesCmd.Flags().String("name", "", "The name of the signal catalog to list information about.")
		iotfleetwise_listSignalCatalogNodesCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
		iotfleetwise_listSignalCatalogNodesCmd.Flags().String("signal-node-type", "", "The type of node in the signal catalog.")
		iotfleetwise_listSignalCatalogNodesCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listSignalCatalogNodesCmd)
}
