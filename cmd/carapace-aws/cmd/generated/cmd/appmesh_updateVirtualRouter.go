package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateVirtualRouterCmd = &cobra.Command{
	Use:   "update-virtual-router",
	Short: "Updates an existing virtual router in a specified service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateVirtualRouterCmd).Standalone()

	appmesh_updateVirtualRouterCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_updateVirtualRouterCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual router resides in.")
	appmesh_updateVirtualRouterCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_updateVirtualRouterCmd.Flags().String("spec", "", "The new virtual router specification to apply.")
	appmesh_updateVirtualRouterCmd.Flags().String("virtual-router-name", "", "The name of the virtual router to update.")
	appmesh_updateVirtualRouterCmd.MarkFlagRequired("mesh-name")
	appmesh_updateVirtualRouterCmd.MarkFlagRequired("spec")
	appmesh_updateVirtualRouterCmd.MarkFlagRequired("virtual-router-name")
	appmeshCmd.AddCommand(appmesh_updateVirtualRouterCmd)
}
