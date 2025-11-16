package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxClusterNodeCmd = &cobra.Command{
	Use:   "delete-kx-cluster-node",
	Short: "Deletes the specified nodes from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxClusterNodeCmd).Standalone()

	finspace_deleteKxClusterNodeCmd.Flags().String("cluster-name", "", "The name of the cluster, for which you want to delete the nodes.")
	finspace_deleteKxClusterNodeCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_deleteKxClusterNodeCmd.Flags().String("node-id", "", "A unique identifier for the node that you want to delete.")
	finspace_deleteKxClusterNodeCmd.MarkFlagRequired("cluster-name")
	finspace_deleteKxClusterNodeCmd.MarkFlagRequired("environment-id")
	finspace_deleteKxClusterNodeCmd.MarkFlagRequired("node-id")
	finspaceCmd.AddCommand(finspace_deleteKxClusterNodeCmd)
}
