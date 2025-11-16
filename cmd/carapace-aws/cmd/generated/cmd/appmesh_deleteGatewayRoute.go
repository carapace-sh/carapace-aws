package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteGatewayRouteCmd = &cobra.Command{
	Use:   "delete-gateway-route",
	Short: "Deletes an existing gateway route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteGatewayRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_deleteGatewayRouteCmd).Standalone()

		appmesh_deleteGatewayRouteCmd.Flags().String("gateway-route-name", "", "The name of the gateway route to delete.")
		appmesh_deleteGatewayRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the gateway route from.")
		appmesh_deleteGatewayRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_deleteGatewayRouteCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to delete the route from.")
		appmesh_deleteGatewayRouteCmd.MarkFlagRequired("gateway-route-name")
		appmesh_deleteGatewayRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_deleteGatewayRouteCmd.MarkFlagRequired("virtual-gateway-name")
	})
	appmeshCmd.AddCommand(appmesh_deleteGatewayRouteCmd)
}
