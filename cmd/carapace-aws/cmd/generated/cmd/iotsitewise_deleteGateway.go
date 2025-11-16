package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteGatewayCmd = &cobra.Command{
	Use:   "delete-gateway",
	Short: "Deletes a gateway from IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteGatewayCmd).Standalone()

		iotsitewise_deleteGatewayCmd.Flags().String("gateway-id", "", "The ID of the gateway to delete.")
		iotsitewise_deleteGatewayCmd.MarkFlagRequired("gateway-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteGatewayCmd)
}
