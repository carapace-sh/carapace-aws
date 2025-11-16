package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteRouteCmd = &cobra.Command{
	Use:   "delete-route",
	Short: "Deletes an existing route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteRouteCmd).Standalone()

	appmesh_deleteRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the route in.")
	appmesh_deleteRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_deleteRouteCmd.Flags().String("route-name", "", "The name of the route to delete.")
	appmesh_deleteRouteCmd.Flags().String("virtual-router-name", "", "The name of the virtual router to delete the route in.")
	appmesh_deleteRouteCmd.MarkFlagRequired("mesh-name")
	appmesh_deleteRouteCmd.MarkFlagRequired("route-name")
	appmesh_deleteRouteCmd.MarkFlagRequired("virtual-router-name")
	appmeshCmd.AddCommand(appmesh_deleteRouteCmd)
}
