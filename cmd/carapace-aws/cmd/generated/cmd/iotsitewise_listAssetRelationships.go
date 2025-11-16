package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetRelationshipsCmd = &cobra.Command{
	Use:   "list-asset-relationships",
	Short: "Retrieves a paginated list of asset relationships for an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetRelationshipsCmd).Standalone()

	iotsitewise_listAssetRelationshipsCmd.Flags().String("asset-id", "", "The ID of the asset.")
	iotsitewise_listAssetRelationshipsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listAssetRelationshipsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listAssetRelationshipsCmd.Flags().String("traversal-type", "", "The type of traversal to use to identify asset relationships.")
	iotsitewise_listAssetRelationshipsCmd.MarkFlagRequired("asset-id")
	iotsitewise_listAssetRelationshipsCmd.MarkFlagRequired("traversal-type")
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetRelationshipsCmd)
}
