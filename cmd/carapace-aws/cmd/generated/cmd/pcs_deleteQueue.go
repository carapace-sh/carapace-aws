package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_deleteQueueCmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "Deletes a job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_deleteQueueCmd).Standalone()

	pcs_deleteQueueCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pcs_deleteQueueCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster of the queue.")
	pcs_deleteQueueCmd.Flags().String("queue-identifier", "", "The name or ID of the queue to delete.")
	pcs_deleteQueueCmd.MarkFlagRequired("cluster-identifier")
	pcs_deleteQueueCmd.MarkFlagRequired("queue-identifier")
	pcsCmd.AddCommand(pcs_deleteQueueCmd)
}
