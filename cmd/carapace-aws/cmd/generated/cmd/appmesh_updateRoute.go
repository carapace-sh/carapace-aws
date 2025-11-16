package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateRouteCmd = &cobra.Command{
	Use:   "update-route",
	Short: "Updates an existing route for a specified service mesh and virtual router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_updateRouteCmd).Standalone()

		appmesh_updateRouteCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_updateRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh that the route resides in.")
		appmesh_updateRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_updateRouteCmd.Flags().String("route-name", "", "The name of the route to update.")
		appmesh_updateRouteCmd.Flags().String("spec", "", "The new route specification to apply.")
		appmesh_updateRouteCmd.Flags().String("virtual-router-name", "", "The name of the virtual router that the route is associated with.")
		appmesh_updateRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_updateRouteCmd.MarkFlagRequired("route-name")
		appmesh_updateRouteCmd.MarkFlagRequired("spec")
		appmesh_updateRouteCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_updateRouteCmd)
}
