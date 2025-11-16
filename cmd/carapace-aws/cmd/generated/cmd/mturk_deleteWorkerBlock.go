package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_deleteWorkerBlockCmd = &cobra.Command{
	Use:   "delete-worker-block",
	Short: "The `DeleteWorkerBlock` operation allows you to reinstate a blocked Worker to work on your HITs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_deleteWorkerBlockCmd).Standalone()

	mturk_deleteWorkerBlockCmd.Flags().String("reason", "", "A message that explains the reason for unblocking the Worker.")
	mturk_deleteWorkerBlockCmd.Flags().String("worker-id", "", "The ID of the Worker to unblock.")
	mturk_deleteWorkerBlockCmd.MarkFlagRequired("worker-id")
	mturkCmd.AddCommand(mturk_deleteWorkerBlockCmd)
}
