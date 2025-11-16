package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listVirtualGatewaysCmd = &cobra.Command{
	Use:   "list-virtual-gateways",
	Short: "Returns a list of existing virtual gateways in a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listVirtualGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listVirtualGatewaysCmd).Standalone()

		appmesh_listVirtualGatewaysCmd.Flags().String("limit", "", "The maximum number of results returned by `ListVirtualGateways` in paginated output.")
		appmesh_listVirtualGatewaysCmd.Flags().String("mesh-name", "", "The name of the service mesh to list virtual gateways in.")
		appmesh_listVirtualGatewaysCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_listVirtualGatewaysCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListVirtualGateways` request where `limit` was used and the results exceeded the value of that parameter.")
		appmesh_listVirtualGatewaysCmd.MarkFlagRequired("mesh-name")
	})
	appmeshCmd.AddCommand(appmesh_listVirtualGatewaysCmd)
}
