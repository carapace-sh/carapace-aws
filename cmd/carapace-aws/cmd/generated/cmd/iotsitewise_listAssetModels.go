package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetModelsCmd = &cobra.Command{
	Use:   "list-asset-models",
	Short: "Retrieves a paginated list of summaries of all asset models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listAssetModelsCmd).Standalone()

		iotsitewise_listAssetModelsCmd.Flags().String("asset-model-types", "", "The type of asset model.")
		iotsitewise_listAssetModelsCmd.Flags().String("asset-model-version", "", "The version alias that specifies the latest or active version of the asset model.")
		iotsitewise_listAssetModelsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listAssetModelsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetModelsCmd)
}
