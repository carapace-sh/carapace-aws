package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_describeMeshCmd = &cobra.Command{
	Use:   "describe-mesh",
	Short: "Describes an existing service mesh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_describeMeshCmd).Standalone()

	appmesh_describeMeshCmd.Flags().String("mesh-name", "", "The name of the service mesh to describe.")
	appmesh_describeMeshCmd.Flags().String("mesh-owner", "", "The Amazon Web Services IAM account ID of the service mesh owner.")
	appmesh_describeMeshCmd.MarkFlagRequired("mesh-name")
	appmeshCmd.AddCommand(appmesh_describeMeshCmd)
}
