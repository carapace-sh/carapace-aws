package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listVirtualNodesCmd = &cobra.Command{
	Use:   "list-virtual-nodes",
	Short: "Returns a list of existing virtual nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listVirtualNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listVirtualNodesCmd).Standalone()

		appmesh_listVirtualNodesCmd.Flags().String("limit", "", "The maximum number of results returned by `ListVirtualNodes` in paginated output.")
		appmesh_listVirtualNodesCmd.Flags().String("mesh-name", "", "The name of the service mesh to list virtual nodes in.")
		appmesh_listVirtualNodesCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_listVirtualNodesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListVirtualNodes` request where `limit` was used and the results exceeded the value of that parameter.")
		appmesh_listVirtualNodesCmd.MarkFlagRequired("mesh-name")
	})
	appmeshCmd.AddCommand(appmesh_listVirtualNodesCmd)
}
