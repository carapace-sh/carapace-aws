package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_updateQueueCmd = &cobra.Command{
	Use:   "update-queue",
	Short: "Updates the compute node group configuration of a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_updateQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_updateQueueCmd).Standalone()

		pcs_updateQueueCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pcs_updateQueueCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster of the queue.")
		pcs_updateQueueCmd.Flags().String("compute-node-group-configurations", "", "The list of compute node group configurations to associate with the queue.")
		pcs_updateQueueCmd.Flags().String("queue-identifier", "", "The name or ID of the queue.")
		pcs_updateQueueCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
		pcs_updateQueueCmd.MarkFlagRequired("cluster-identifier")
		pcs_updateQueueCmd.MarkFlagRequired("queue-identifier")
	})
	pcsCmd.AddCommand(pcs_updateQueueCmd)
}
