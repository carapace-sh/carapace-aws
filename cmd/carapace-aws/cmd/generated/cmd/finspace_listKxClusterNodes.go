package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxClusterNodesCmd = &cobra.Command{
	Use:   "list-kx-cluster-nodes",
	Short: "Lists all the nodes in a kdb cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxClusterNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_listKxClusterNodesCmd).Standalone()

		finspace_listKxClusterNodesCmd.Flags().String("cluster-name", "", "A unique name for the cluster.")
		finspace_listKxClusterNodesCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_listKxClusterNodesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		finspace_listKxClusterNodesCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
		finspace_listKxClusterNodesCmd.MarkFlagRequired("cluster-name")
		finspace_listKxClusterNodesCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_listKxClusterNodesCmd)
}
