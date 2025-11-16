package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateGatewayCmd = &cobra.Command{
	Use:   "update-gateway",
	Short: "Updates a gateway's name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateGatewayCmd).Standalone()

	iotsitewise_updateGatewayCmd.Flags().String("gateway-id", "", "The ID of the gateway to update.")
	iotsitewise_updateGatewayCmd.Flags().String("gateway-name", "", "A unique name for the gateway.")
	iotsitewise_updateGatewayCmd.MarkFlagRequired("gateway-id")
	iotsitewise_updateGatewayCmd.MarkFlagRequired("gateway-name")
	iotsitewiseCmd.AddCommand(iotsitewise_updateGatewayCmd)
}
