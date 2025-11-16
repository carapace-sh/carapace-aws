package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_createWorkerBlockCmd = &cobra.Command{
	Use:   "create-worker-block",
	Short: "The `CreateWorkerBlock` operation allows you to prevent a Worker from working on your HITs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_createWorkerBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_createWorkerBlockCmd).Standalone()

		mturk_createWorkerBlockCmd.Flags().String("reason", "", "A message explaining the reason for blocking the Worker.")
		mturk_createWorkerBlockCmd.Flags().String("worker-id", "", "The ID of the Worker to block.")
		mturk_createWorkerBlockCmd.MarkFlagRequired("reason")
		mturk_createWorkerBlockCmd.MarkFlagRequired("worker-id")
	})
	mturkCmd.AddCommand(mturk_createWorkerBlockCmd)
}
