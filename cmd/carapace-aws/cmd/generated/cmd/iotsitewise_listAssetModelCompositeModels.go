package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetModelCompositeModelsCmd = &cobra.Command{
	Use:   "list-asset-model-composite-models",
	Short: "Retrieves a paginated list of composite models associated with the asset model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetModelCompositeModelsCmd).Standalone()

	iotsitewise_listAssetModelCompositeModelsCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
	iotsitewise_listAssetModelCompositeModelsCmd.Flags().String("asset-model-version", "", "The version alias that specifies the latest or active version of the asset model.")
	iotsitewise_listAssetModelCompositeModelsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listAssetModelCompositeModelsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listAssetModelCompositeModelsCmd.MarkFlagRequired("asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetModelCompositeModelsCmd)
}
