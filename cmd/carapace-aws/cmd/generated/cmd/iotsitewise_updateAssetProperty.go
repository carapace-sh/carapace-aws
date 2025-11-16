package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateAssetPropertyCmd = &cobra.Command{
	Use:   "update-asset-property",
	Short: "Updates an asset property's alias and notification state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateAssetPropertyCmd).Standalone()

	iotsitewise_updateAssetPropertyCmd.Flags().String("asset-id", "", "The ID of the asset to be updated.")
	iotsitewise_updateAssetPropertyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateAssetPropertyCmd.Flags().String("property-alias", "", "The alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature`).")
	iotsitewise_updateAssetPropertyCmd.Flags().String("property-id", "", "The ID of the asset property to be updated.")
	iotsitewise_updateAssetPropertyCmd.Flags().String("property-notification-state", "", "The MQTT notification state (enabled or disabled) for this asset property.")
	iotsitewise_updateAssetPropertyCmd.Flags().String("property-unit", "", "The unit of measure (such as Newtons or RPM) of the asset property.")
	iotsitewise_updateAssetPropertyCmd.MarkFlagRequired("asset-id")
	iotsitewise_updateAssetPropertyCmd.MarkFlagRequired("property-id")
	iotsitewiseCmd.AddCommand(iotsitewise_updateAssetPropertyCmd)
}
