package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listVirtualServicesCmd = &cobra.Command{
	Use:   "list-virtual-services",
	Short: "Returns a list of existing virtual services in a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listVirtualServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listVirtualServicesCmd).Standalone()

		appmesh_listVirtualServicesCmd.Flags().String("limit", "", "The maximum number of results returned by `ListVirtualServices` in paginated output.")
		appmesh_listVirtualServicesCmd.Flags().String("mesh-name", "", "The name of the service mesh to list virtual services in.")
		appmesh_listVirtualServicesCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_listVirtualServicesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListVirtualServices` request where `limit` was used and the results exceeded the value of that parameter.")
		appmesh_listVirtualServicesCmd.MarkFlagRequired("mesh-name")
	})
	appmeshCmd.AddCommand(appmesh_listVirtualServicesCmd)
}
