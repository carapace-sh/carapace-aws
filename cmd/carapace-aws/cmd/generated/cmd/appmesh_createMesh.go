package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_createMeshCmd = &cobra.Command{
	Use:   "create-mesh",
	Short: "Creates a service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_createMeshCmd).Standalone()

	appmesh_createMeshCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appmesh_createMeshCmd.Flags().String("mesh-name", "", "The name to use for the service mesh.")
	appmesh_createMeshCmd.Flags().String("spec", "", "The service mesh specification to apply.")
	appmesh_createMeshCmd.Flags().String("tags", "", "Optional metadata that you can apply to the service mesh to assist with categorization and organization.")
	appmesh_createMeshCmd.MarkFlagRequired("mesh-name")
	appmeshCmd.AddCommand(appmesh_createMeshCmd)
}
