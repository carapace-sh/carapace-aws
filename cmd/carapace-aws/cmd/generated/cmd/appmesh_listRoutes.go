package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listRoutesCmd = &cobra.Command{
	Use:   "list-routes",
	Short: "Returns a list of existing routes in a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listRoutesCmd).Standalone()

		appmesh_listRoutesCmd.Flags().String("limit", "", "The maximum number of results returned by `ListRoutes` in paginated output.")
		appmesh_listRoutesCmd.Flags().String("mesh-name", "", "The name of the service mesh to list routes in.")
		appmesh_listRoutesCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_listRoutesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListRoutes` request where `limit` was used and the results exceeded the value of that parameter.")
		appmesh_listRoutesCmd.Flags().String("virtual-router-name", "", "The name of the virtual router to list routes in.")
		appmesh_listRoutesCmd.MarkFlagRequired("mesh-name")
		appmesh_listRoutesCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_listRoutesCmd)
}
