package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteVirtualNodeCmd = &cobra.Command{
	Use:   "delete-virtual-node",
	Short: "Deletes an existing virtual node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteVirtualNodeCmd).Standalone()

	appmesh_deleteVirtualNodeCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the virtual node in.")
	appmesh_deleteVirtualNodeCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_deleteVirtualNodeCmd.Flags().String("virtual-node-name", "", "The name of the virtual node to delete.")
	appmesh_deleteVirtualNodeCmd.MarkFlagRequired("mesh-name")
	appmesh_deleteVirtualNodeCmd.MarkFlagRequired("virtual-node-name")
	appmeshCmd.AddCommand(appmesh_deleteVirtualNodeCmd)
}
