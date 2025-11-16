package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createRouteCmd = &cobra.Command{
	Use:   "create-route",
	Short: "Creates a route that is associated with a virtual router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_createRouteCmd).Standalone()

		appmesh_createRouteCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_createRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the route in.")
		appmesh_createRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_createRouteCmd.Flags().String("route-name", "", "The name to use for the route.")
		appmesh_createRouteCmd.Flags().String("spec", "", "The route specification to apply.")
		appmesh_createRouteCmd.Flags().String("tags", "", "Optional metadata that you can apply to the route to assist with categorization and organization.")
		appmesh_createRouteCmd.Flags().String("virtual-router-name", "", "The name of the virtual router in which to create the route.")
		appmesh_createRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_createRouteCmd.MarkFlagRequired("route-name")
		appmesh_createRouteCmd.MarkFlagRequired("spec")
		appmesh_createRouteCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_createRouteCmd)
}
