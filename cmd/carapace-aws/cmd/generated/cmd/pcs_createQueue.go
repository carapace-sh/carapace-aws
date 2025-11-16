package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_createQueueCmd = &cobra.Command{
	Use:   "create-queue",
	Short: "Creates a job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_createQueueCmd).Standalone()

	pcs_createQueueCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pcs_createQueueCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster for which to create a queue.")
	pcs_createQueueCmd.Flags().String("compute-node-group-configurations", "", "The list of compute node group configurations to associate with the queue.")
	pcs_createQueueCmd.Flags().String("queue-name", "", "A name to identify the queue.")
	pcs_createQueueCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
	pcs_createQueueCmd.Flags().String("tags", "", "1 or more tags added to the resource.")
	pcs_createQueueCmd.MarkFlagRequired("cluster-identifier")
	pcs_createQueueCmd.MarkFlagRequired("queue-name")
	pcsCmd.AddCommand(pcs_createQueueCmd)
}
