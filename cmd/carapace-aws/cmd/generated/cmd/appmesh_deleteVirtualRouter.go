package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteVirtualRouterCmd = &cobra.Command{
	Use:   "delete-virtual-router",
	Short: "Deletes an existing virtual router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteVirtualRouterCmd).Standalone()

	appmesh_deleteVirtualRouterCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the virtual router in.")
	appmesh_deleteVirtualRouterCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_deleteVirtualRouterCmd.Flags().String("virtual-router-name", "", "The name of the virtual router to delete.")
	appmesh_deleteVirtualRouterCmd.MarkFlagRequired("mesh-name")
	appmesh_deleteVirtualRouterCmd.MarkFlagRequired("virtual-router-name")
	appmeshCmd.AddCommand(appmesh_deleteVirtualRouterCmd)
}
