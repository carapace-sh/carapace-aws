package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteVirtualGatewayCmd = &cobra.Command{
	Use:   "delete-virtual-gateway",
	Short: "Deletes an existing virtual gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteVirtualGatewayCmd).Standalone()

	appmesh_deleteVirtualGatewayCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the virtual gateway from.")
	appmesh_deleteVirtualGatewayCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_deleteVirtualGatewayCmd.Flags().String("virtual-gateway-name", "", "The name of the virtual gateway to delete.")
	appmesh_deleteVirtualGatewayCmd.MarkFlagRequired("mesh-name")
	appmesh_deleteVirtualGatewayCmd.MarkFlagRequired("virtual-gateway-name")
	appmeshCmd.AddCommand(appmesh_deleteVirtualGatewayCmd)
}
