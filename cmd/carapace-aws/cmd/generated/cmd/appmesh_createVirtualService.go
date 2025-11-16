package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createVirtualServiceCmd = &cobra.Command{
	Use:   "create-virtual-service",
	Short: "Creates a virtual service within a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createVirtualServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_createVirtualServiceCmd).Standalone()

		appmesh_createVirtualServiceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appmesh_createVirtualServiceCmd.Flags().String("mesh-name", "", "The name of the service mesh to create the virtual service in.")
		appmesh_createVirtualServiceCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
		appmesh_createVirtualServiceCmd.Flags().String("spec", "", "The virtual service specification to apply.")
		appmesh_createVirtualServiceCmd.Flags().String("tags", "", "Optional metadata that you can apply to the virtual service to assist with categorization and organization.")
		appmesh_createVirtualServiceCmd.Flags().String("virtual-service-name", "", "The name to use for the virtual service.")
		appmesh_createVirtualServiceCmd.MarkFlagRequired("mesh-name")
		appmesh_createVirtualServiceCmd.MarkFlagRequired("spec")
		appmesh_createVirtualServiceCmd.MarkFlagRequired("virtual-service-name")
	})
	appmeshCmd.AddCommand(appmesh_createVirtualServiceCmd)
}
