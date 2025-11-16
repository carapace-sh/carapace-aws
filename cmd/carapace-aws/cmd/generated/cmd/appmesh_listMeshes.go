package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listMeshesCmd = &cobra.Command{
	Use:   "list-meshes",
	Short: "Returns a list of existing service meshes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listMeshesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmesh_listMeshesCmd).Standalone()

		appmesh_listMeshesCmd.Flags().String("limit", "", "The maximum number of results returned by `ListMeshes` in paginated output.")
		appmesh_listMeshesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListMeshes` request where `limit` was used and the results exceeded the value of that parameter.")
	})
	appmeshCmd.AddCommand(appmesh_listMeshesCmd)
}
