package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listModelManifestNodesCmd = &cobra.Command{
	Use:   "list-model-manifest-nodes",
	Short: "Lists information about nodes specified in a vehicle model (model manifest).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listModelManifestNodesCmd).Standalone()

	iotfleetwise_listModelManifestNodesCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
	iotfleetwise_listModelManifestNodesCmd.Flags().String("name", "", "The name of the vehicle model to list information about.")
	iotfleetwise_listModelManifestNodesCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	iotfleetwise_listModelManifestNodesCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_listModelManifestNodesCmd)
}
