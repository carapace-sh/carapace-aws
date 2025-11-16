package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetModelPropertiesCmd = &cobra.Command{
	Use:   "list-asset-model-properties",
	Short: "Retrieves a paginated list of properties associated with an asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetModelPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listAssetModelPropertiesCmd).Standalone()

		iotsitewise_listAssetModelPropertiesCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
		iotsitewise_listAssetModelPropertiesCmd.Flags().String("asset-model-version", "", "The version alias that specifies the latest or active version of the asset model.")
		iotsitewise_listAssetModelPropertiesCmd.Flags().String("filter", "", "Filters the requested list of asset model properties.")
		iotsitewise_listAssetModelPropertiesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listAssetModelPropertiesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listAssetModelPropertiesCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetModelPropertiesCmd)
}
