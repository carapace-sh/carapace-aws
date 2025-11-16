package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createVirtualNodeCmd = &cobra.Command{
	Use:   "create-virtual-node",
	Short: "Creates a virtual node within a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createVirtualNodeCmd).Standalone()

	appmesh_createVirtualNodeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_createVirtualNodeCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the virtual node in.")
	appmesh_createVirtualNodeCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_createVirtualNodeCmd.Flags().String("spec", "", "The virtual node specification to apply.")
	appmesh_createVirtualNodeCmd.Flags().String("tags", "", "Optional metadata that you can apply to the virtual node to assist with categorization and organization.")
	appmesh_createVirtualNodeCmd.Flags().String("virtual-node-name", "", "The name to use for the virtual node.")
	appmesh_createVirtualNodeCmd.MarkFlagRequired("mesh-name")
	appmesh_createVirtualNodeCmd.MarkFlagRequired("spec")
	appmesh_createVirtualNodeCmd.MarkFlagRequired("virtual-node-name")
	appmeshCmd.AddCommand(appmesh_createVirtualNodeCmd)
}
