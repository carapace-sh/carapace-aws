package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_getComputeNodeGroupCmd = &cobra.Command{
	Use:   "get-compute-node-group",
	Short: "Returns detailed information about a compute node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_getComputeNodeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_getComputeNodeGroupCmd).Standalone()

		pcs_getComputeNodeGroupCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster.")
		pcs_getComputeNodeGroupCmd.Flags().String("compute-node-group-identifier", "", "The name or ID of the compute node group.")
		pcs_getComputeNodeGroupCmd.MarkFlagRequired("cluster-identifier")
		pcs_getComputeNodeGroupCmd.MarkFlagRequired("compute-node-group-identifier")
	})
	pcsCmd.AddCommand(pcs_getComputeNodeGroupCmd)
}
