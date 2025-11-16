package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listAssetPropertiesCmd = &cobra.Command{
	Use:   "list-asset-properties",
	Short: "Retrieves a paginated list of properties associated with an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listAssetPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listAssetPropertiesCmd).Standalone()

		iotsitewise_listAssetPropertiesCmd.Flags().String("asset-id", "", "The ID of the asset.")
		iotsitewise_listAssetPropertiesCmd.Flags().String("filter", "", "Filters the requested list of asset properties.")
		iotsitewise_listAssetPropertiesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listAssetPropertiesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listAssetPropertiesCmd.MarkFlagRequired("asset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listAssetPropertiesCmd)
}
