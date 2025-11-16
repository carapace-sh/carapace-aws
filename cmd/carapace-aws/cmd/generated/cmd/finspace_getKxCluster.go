package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxClusterCmd = &cobra.Command{
	Use:   "get-kx-cluster",
	Short: "Retrieves information about a kdb cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxClusterCmd).Standalone()

	finspace_getKxClusterCmd.Flags().String("cluster-name", "", "The name of the cluster that you want to retrieve.")
	finspace_getKxClusterCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_getKxClusterCmd.MarkFlagRequired("cluster-name")
	finspace_getKxClusterCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_getKxClusterCmd)
}
