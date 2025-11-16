package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetsCmd = &cobra.Command{
	Use:   "list-assets",
	Short: "Retrieves a paginated list of asset summaries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listAssetsCmd).Standalone()

		iotsitewise_listAssetsCmd.Flags().String("asset-model-id", "", "The ID of the asset model by which to filter the list of assets.")
		iotsitewise_listAssetsCmd.Flags().String("filter", "", "The filter for the requested list of assets.")
		iotsitewise_listAssetsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listAssetsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetsCmd)
}
