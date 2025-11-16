package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_getQueueCmd = &cobra.Command{
	Use:   "get-queue",
	Short: "Returns detailed information about a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_getQueueCmd).Standalone()

	pcs_getQueueCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster of the queue.")
	pcs_getQueueCmd.Flags().String("queue-identifier", "", "The name or ID of the queue.")
	pcs_getQueueCmd.MarkFlagRequired("cluster-identifier")
	pcs_getQueueCmd.MarkFlagRequired("queue-identifier")
	pcsCmd.AddCommand(pcs_getQueueCmd)
}
