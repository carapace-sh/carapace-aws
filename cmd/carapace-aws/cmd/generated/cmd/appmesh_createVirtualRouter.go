package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createVirtualRouterCmd = &cobra.Command{
	Use:   "create-virtual-router",
	Short: "Creates a virtual router within a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createVirtualRouterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_createVirtualRouterCmd).Standalone()

		appmesh_createVirtualRouterCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_createVirtualRouterCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the virtual router in.")
		appmesh_createVirtualRouterCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_createVirtualRouterCmd.Flags().String("spec", "", "The virtual router specification to apply.")
		appmesh_createVirtualRouterCmd.Flags().String("tags", "", "Optional metadata that you can apply to the virtual router to assist with categorization and organization.")
		appmesh_createVirtualRouterCmd.Flags().String("virtual-router-name", "", "The name to use for the virtual router.")
		appmesh_createVirtualRouterCmd.MarkFlagRequired("mesh-name")
		appmesh_createVirtualRouterCmd.MarkFlagRequired("spec")
		appmesh_createVirtualRouterCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_createVirtualRouterCmd)
}
