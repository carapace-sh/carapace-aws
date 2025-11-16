package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_deleteMeshCmd = &cobra.Command{
	Use:   "delete-mesh",
	Short: "Deletes an existing service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_deleteMeshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_deleteMeshCmd).Standalone()

		appmesh_deleteMeshCmd.Flags().String("mesh-name", "", "The name of the service mesh to delete.")
		appmesh_deleteMeshCmd.MarkFlagRequired("mesh-name")
	})
	appmeshCmd.AddCommand(appmesh_deleteMeshCmd)
}
