package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_deleteComputeNodeGroupCmd = &cobra.Command{
	Use:   "delete-compute-node-group",
	Short: "Deletes a compute node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_deleteComputeNodeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_deleteComputeNodeGroupCmd).Standalone()

		pcs_deleteComputeNodeGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pcs_deleteComputeNodeGroupCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster of the compute node group.")
		pcs_deleteComputeNodeGroupCmd.Flags().String("compute-node-group-identifier", "", "The name or ID of the compute node group to delete.")
		pcs_deleteComputeNodeGroupCmd.MarkFlagRequired("cluster-identifier")
		pcs_deleteComputeNodeGroupCmd.MarkFlagRequired("compute-node-group-identifier")
	})
	pcsCmd.AddCommand(pcs_deleteComputeNodeGroupCmd)
}
