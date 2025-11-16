package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_updateMeshCmd = &cobra.Command{
	Use:   "update-mesh",
	Short: "Updates an existing service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_updateMeshCmd).Standalone()

	appmesh_updateMeshCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_updateMeshCmd.Flags().String("mesh-name", "", "The name of the service mesh to update.")
	appmesh_updateMeshCmd.Flags().String("spec", "", "The service mesh specification to apply.")
	appmesh_updateMeshCmd.MarkFlagRequired("mesh-name")
	appmeshCmd.AddCommand(appmesh_updateMeshCmd)
}
