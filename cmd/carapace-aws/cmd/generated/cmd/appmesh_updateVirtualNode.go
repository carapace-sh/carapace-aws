package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateVirtualNodeCmd = &cobra.Command{
	Use:   "update-virtual-node",
	Short: "Updates an existing virtual node in a specified service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateVirtualNodeCmd).Standalone()

	appmesh_updateVirtualNodeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_updateVirtualNodeCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual node resides in.")
	appmesh_updateVirtualNodeCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_updateVirtualNodeCmd.Flags().String("spec", "", "The new virtual node specification to apply.")
	appmesh_updateVirtualNodeCmd.Flags().String("virtual-node-name", "", "The name of the virtual node to update.")
	appmesh_updateVirtualNodeCmd.MarkFlagRequired("mesh-name")
	appmesh_updateVirtualNodeCmd.MarkFlagRequired("spec")
	appmesh_updateVirtualNodeCmd.MarkFlagRequired("virtual-node-name")
	appmeshCmd.AddCommand(appmesh_updateVirtualNodeCmd)
}
