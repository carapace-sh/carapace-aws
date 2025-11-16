package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateVirtualGatewayCmd = &cobra.Command{
	Use:   "update-virtual-gateway",
	Short: "Updates an existing virtual gateway in a specified service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateVirtualGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_updateVirtualGatewayCmd).Standalone()

		appmesh_updateVirtualGatewayCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_updateVirtualGatewayCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual gateway resides in.")
		appmesh_updateVirtualGatewayCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_updateVirtualGatewayCmd.Flags().String("spec", "", "The new virtual gateway specification to apply.")
		appmesh_updateVirtualGatewayCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to update.")
		appmesh_updateVirtualGatewayCmd.MarkFlagRequired("mesh-name")
		appmesh_updateVirtualGatewayCmd.MarkFlagRequired("spec")
		appmesh_updateVirtualGatewayCmd.MarkFlagRequired("virtual-gateway-name")
	})
	appmeshCmd.AddCommand(appmesh_updateVirtualGatewayCmd)
}
