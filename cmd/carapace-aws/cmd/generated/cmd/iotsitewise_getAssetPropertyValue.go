package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_getAssetPropertyValueCmd = &cobra.Command{
	Use:   "get-asset-property-value",
	Short: "Gets an asset property's current value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_getAssetPropertyValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_getAssetPropertyValueCmd).Standalone()

		iotsitewise_getAssetPropertyValueCmd.Flags().String("asset-id", "", "The ID of the asset, in UUID format.")
		iotsitewise_getAssetPropertyValueCmd.Flags().String("property-alias", "", "The alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature`).")
		iotsitewise_getAssetPropertyValueCmd.Flags().String("property-id", "", "The ID of the asset property, in UUID format.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_getAssetPropertyValueCmd)
}
