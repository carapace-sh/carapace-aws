package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createGatewayCmd = &cobra.Command{
	Use:   "create-gateway",
	Short: "Creates a gateway, which is a virtual or edge device that delivers industrial data streams from local servers to IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createGatewayCmd).Standalone()

	iotsitewise_createGatewayCmd.Flags().String("gateway-name", "", "A unique name for the gateway.")
	iotsitewise_createGatewayCmd.Flags().String("gateway-platform", "", "The gateway's platform.")
	iotsitewise_createGatewayCmd.Flags().String("gateway-version", "", "The version of the gateway to create.")
	iotsitewise_createGatewayCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the gateway.")
	iotsitewise_createGatewayCmd.MarkFlagRequired("gateway-name")
	iotsitewise_createGatewayCmd.MarkFlagRequired("gateway-platform")
	iotsitewiseCmd.AddCommand(iotsitewise_createGatewayCmd)
}
