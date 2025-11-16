package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeRouteCmd = &cobra.Command{
	Use:   "describe-route",
	Short: "Describes an existing route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_describeRouteCmd).Standalone()

		appmesh_describeRouteCmd.Flags().String("mesh-name", "", "The name of the service mesh that the route resides in.")
		appmesh_describeRouteCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_describeRouteCmd.Flags().String("route-name", "", "The name of the route to describe.")
		appmesh_describeRouteCmd.Flags().String("virtual-router-name", "", "The name of the virtual router that the route is associated with.")
		appmesh_describeRouteCmd.MarkFlagRequired("mesh-name")
		appmesh_describeRouteCmd.MarkFlagRequired("route-name")
		appmesh_describeRouteCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_describeRouteCmd)
}
