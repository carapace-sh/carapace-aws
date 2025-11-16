package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateVirtualServiceCmd = &cobra.Command{
	Use:   "update-virtual-service",
	Short: "Updates an existing virtual service in a specified service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateVirtualServiceCmd).Standalone()

	appmesh_updateVirtualServiceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_updateVirtualServiceCmd.Flags().String("mesh-name", "", "The name of the service mesh that the virtual service resides in.")
	appmesh_updateVirtualServiceCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_updateVirtualServiceCmd.Flags().String("spec", "", "The new virtual service specification to apply.")
	appmesh_updateVirtualServiceCmd.Flags().String("virtual-service-name", "", "The name of the virtual service to update.")
	appmesh_updateVirtualServiceCmd.MarkFlagRequired("mesh-name")
	appmesh_updateVirtualServiceCmd.MarkFlagRequired("spec")
	appmesh_updateVirtualServiceCmd.MarkFlagRequired("virtual-service-name")
	appmeshCmd.AddCommand(appmesh_updateVirtualServiceCmd)
}
