package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_listComputeNodeGroupsCmd = &cobra.Command{
	Use:   "list-compute-node-groups",
	Short: "Returns a list of all compute node groups associated with a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_listComputeNodeGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_listComputeNodeGroupsCmd).Standalone()

		pcs_listComputeNodeGroupsCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to list compute node groups for.")
		pcs_listComputeNodeGroupsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		pcs_listComputeNodeGroupsCmd.Flags().String("next-token", "", "The value of `nextToken` is a unique pagination token for each page of results returned.")
		pcs_listComputeNodeGroupsCmd.MarkFlagRequired("cluster-identifier")
	})
	pcsCmd.AddCommand(pcs_listComputeNodeGroupsCmd)
}
