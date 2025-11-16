package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateGatewayRouteCmd = &cobra.Command{
	Use:   "update-gateway-route",
	Short: "Updates an existing gateway route that is associated to a specified virtual gateway in a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateGatewayRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_updateGatewayRouteCmd).Standalone()

		appmesh_updateGatewayRouteCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_updateGatewayRouteCmd.Flags().String("gateway-route-name", "", "The name of the gateway route to update.")
		appmesh_updateGatewayRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh that the gateway route resides in.")
		appmesh_updateGatewayRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_updateGatewayRouteCmd.Flags().String("spec", "", "The new gateway route specification to apply.")
		appmesh_updateGatewayRouteCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway that the gateway route is associated with.")
		appmesh_updateGatewayRouteCmd.MarkFlagRequired("gateway-route-name")
		appmesh_updateGatewayRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_updateGatewayRouteCmd.MarkFlagRequired("spec")
		appmesh_updateGatewayRouteCmd.MarkFlagRequired("virtual-gateway-name")
	})
	appmeshCmd.AddCommand(appmesh_updateGatewayRouteCmd)
}
