package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeVirtualNodeCmd = &cobra.Command{
	Use:   "describe-virtual-node",
	Short: "Describes an existing virtual node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeVirtualNodeCmd).Standalone()

	appmesh_describeVirtualNodeCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual node resides in.")
	appmesh_describeVirtualNodeCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_describeVirtualNodeCmd.Flags().String("virtual-node-name", "", "The name of the virtual node to describe.")
	appmesh_describeVirtualNodeCmd.MarkFlagRequired("mesh-name")
	appmesh_describeVirtualNodeCmd.MarkFlagRequired("virtual-node-name")
	appmeshCmd.AddCommand(appmesh_describeVirtualNodeCmd)
}
