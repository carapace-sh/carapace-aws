package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeVirtualServiceCmd = &cobra.Command{
	Use:   "describe-virtual-service",
	Short: "Describes an existing virtual service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeVirtualServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_describeVirtualServiceCmd).Standalone()

		appmesh_describeVirtualServiceCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual service resides in.")
		appmesh_describeVirtualServiceCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_describeVirtualServiceCmd.Flags().String("virtual-service-name", "", "The name of the virtual service to describe.")
		appmesh_describeVirtualServiceCmd.MarkFlagRequired("mesh-name")
		appmesh_describeVirtualServiceCmd.MarkFlagRequired("virtual-service-name")
	})
	appmeshCmd.AddCommand(appmesh_describeVirtualServiceCmd)
}
