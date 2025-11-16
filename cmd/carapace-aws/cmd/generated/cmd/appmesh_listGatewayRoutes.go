package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listGatewayRoutesCmd = &cobra.Command{
	Use:   "list-gateway-routes",
	Short: "Returns a list of existing gateway routes that are associated to a virtual gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listGatewayRoutesCmd).Standalone()

	appmesh_listGatewayRoutesCmd.Flags().String("limit", "", "The maximum number of results returned by `ListGatewayRoutes` in paginated output.")
	appmesh_listGatewayRoutesCmd.Flags().String("mesh-name", "", "The name of the service mesh to list gateway routes in.")
	appmesh_listGatewayRoutesCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_listGatewayRoutesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListGatewayRoutes` request where `limit` was used and the results exceeded the value of that parameter.")
	appmesh_listGatewayRoutesCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to list gateway routes in.")
	appmesh_listGatewayRoutesCmd.MarkFlagRequired("mesh-name")
	appmesh_listGatewayRoutesCmd.MarkFlagRequired("virtual-gateway-name")
	appmeshCmd.AddCommand(appmesh_listGatewayRoutesCmd)
}
