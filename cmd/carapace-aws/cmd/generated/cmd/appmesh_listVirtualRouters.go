package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listVirtualRoutersCmd = &cobra.Command{
	Use:   "list-virtual-routers",
	Short: "Returns a list of existing virtual routers in a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listVirtualRoutersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listVirtualRoutersCmd).Standalone()

		appmesh_listVirtualRoutersCmd.Flags().String("limit", "", "The maximum number of results returned by `ListVirtualRouters` in paginated output.")
		appmesh_listVirtualRoutersCmd.Flags().String("mesh-name", "", "The name of the service mesh to list virtual routers in.")
		appmesh_listVirtualRoutersCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_listVirtualRoutersCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListVirtualRouters` request where `limit` was used and the results exceeded the value of that parameter.")
		appmesh_listVirtualRoutersCmd.MarkFlagRequired("mesh-name")
	})
	appmeshCmd.AddCommand(appmesh_listVirtualRoutersCmd)
}
