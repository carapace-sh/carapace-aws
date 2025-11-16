package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createGatewayRouteCmd = &cobra.Command{
	Use:   "create-gateway-route",
	Short: "Creates a gateway route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createGatewayRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_createGatewayRouteCmd).Standalone()

		appmesh_createGatewayRouteCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_createGatewayRouteCmd.Flags().String("gateway-route-name", "", "The name to use for the gateway route.")
		appmesh_createGatewayRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the gateway route in.")
		appmesh_createGatewayRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_createGatewayRouteCmd.Flags().String("spec", "", "The gateway route specification to apply.")
		appmesh_createGatewayRouteCmd.Flags().String("tags", "", "Optional metadata that you can apply to the gateway route to assist with categorization and organization.")
		appmesh_createGatewayRouteCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to associate the gateway route with.")
		appmesh_createGatewayRouteCmd.MarkFlagRequired("gateway-route-name")
		appmesh_createGatewayRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_createGatewayRouteCmd.MarkFlagRequired("spec")
		appmesh_createGatewayRouteCmd.MarkFlagRequired("virtual-gateway-name")
	})
	appmeshCmd.AddCommand(appmesh_createGatewayRouteCmd)
}
