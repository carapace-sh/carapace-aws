package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeVirtualGatewayCmd = &cobra.Command{
	Use:   "describe-virtual-gateway",
	Short: "Describes an existing virtual gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeVirtualGatewayCmd).Standalone()

	appmesh_describeVirtualGatewayCmd.Flags().String("mesh-name", "", "The name of the service mesh that the gateway route resides in.")
	appmesh_describeVirtualGatewayCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_describeVirtualGatewayCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to describe.")
	appmesh_describeVirtualGatewayCmd.MarkFlagRequired("mesh-name")
	appmesh_describeVirtualGatewayCmd.MarkFlagRequired("virtual-gateway-name")
	appmeshCmd.AddCommand(appmesh_describeVirtualGatewayCmd)
}
