package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteVirtualServiceCmd = &cobra.Command{
	Use:   "delete-virtual-service",
	Short: "Deletes an existing virtual service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteVirtualServiceCmd).Standalone()

	appmesh_deleteVirtualServiceCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete the virtual service in.")
	appmesh_deleteVirtualServiceCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_deleteVirtualServiceCmd.Flags().String("virtual-service-name", "", "The name of the virtual service to delete.")
	appmesh_deleteVirtualServiceCmd.MarkFlagRequired("mesh-name")
	appmesh_deleteVirtualServiceCmd.MarkFlagRequired("virtual-service-name")
	appmeshCmd.AddCommand(appmesh_deleteVirtualServiceCmd)
}
