package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssociatedAssetsCmd = &cobra.Command{
	Use:   "list-associated-assets",
	Short: "Retrieves a paginated list of associated assets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssociatedAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listAssociatedAssetsCmd).Standalone()

		iotsitewise_listAssociatedAssetsCmd.Flags().String("asset-id", "", "The ID of the asset to query.")
		iotsitewise_listAssociatedAssetsCmd.Flags().String("hierarchy-id", "", "(Optional) If you don't provide a `hierarchyId`, all the immediate assets in the `traversalDirection` will be returned.")
		iotsitewise_listAssociatedAssetsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listAssociatedAssetsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listAssociatedAssetsCmd.Flags().String("traversal-direction", "", "The direction to list associated assets.")
		iotsitewise_listAssociatedAssetsCmd.MarkFlagRequired("asset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listAssociatedAssetsCmd)
}
