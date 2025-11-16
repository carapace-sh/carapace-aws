package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeVirtualRouterCmd = &cobra.Command{
	Use:   "describe-virtual-router",
	Short: "Describes an existing virtual router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeVirtualRouterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_describeVirtualRouterCmd).Standalone()

		appmesh_describeVirtualRouterCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual router resides in.")
		appmesh_describeVirtualRouterCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_describeVirtualRouterCmd.Flags().String("virtual-router-name", "", "The name of the virtual router to describe.")
		appmesh_describeVirtualRouterCmd.MarkFlagRequired("mesh-name")
		appmesh_describeVirtualRouterCmd.MarkFlagRequired("virtual-router-name")
	})
	appmeshCmd.AddCommand(appmesh_describeVirtualRouterCmd)
}
