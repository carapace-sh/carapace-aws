package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeGatewayRouteCmd = &cobra.Command{
	Use:   "describe-gateway-route",
	Short: "Describes an existing gateway route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeGatewayRouteCmd).Standalone()

	appmesh_describeGatewayRouteCmd.Flags().String("gateway-route-name", "", "The name of the gateway route to describe.")
	appmesh_describeGatewayRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh that the gateway route resides in.")
	appmesh_describeGatewayRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_describeGatewayRouteCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway that the gateway route is associated with.")
	appmesh_describeGatewayRouteCmd.MarkFlagRequired("gateway-route-name")
	appmesh_describeGatewayRouteCmd.MarkFlagRequired("mesh-name")
	appmesh_describeGatewayRouteCmd.MarkFlagRequired("virtual-gateway-name")
	appmeshCmd.AddCommand(appmesh_describeGatewayRouteCmd)
}
