package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createVirtualGatewayCmd = &cobra.Command{
	Use:   "create-virtual-gateway",
	Short: "Creates a virtual gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createVirtualGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_createVirtualGatewayCmd).Standalone()

		appmesh_createVirtualGatewayCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_createVirtualGatewayCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the virtual gateway in.")
		appmesh_createVirtualGatewayCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_createVirtualGatewayCmd.Flags().String("spec", "", "The virtual gateway specification to apply.")
		appmesh_createVirtualGatewayCmd.Flags().String("tags", "", "Optional metadata that you can apply to the virtual gateway to assist with categorization and organization.")
		appmesh_createVirtualGatewayCmd.Flags().String("virtual-gateway-name", "", "The name to use for the virtual gateway.")
		appmesh_createVirtualGatewayCmd.MarkFlagRequired("mesh-name")
		appmesh_createVirtualGatewayCmd.MarkFlagRequired("spec")
		appmesh_createVirtualGatewayCmd.MarkFlagRequired("virtual-gateway-name")
	})
	appmeshCmd.AddCommand(appmesh_createVirtualGatewayCmd)
}
