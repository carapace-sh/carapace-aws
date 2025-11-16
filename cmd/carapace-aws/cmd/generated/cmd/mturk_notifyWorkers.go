package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_notifyWorkersCmd = &cobra.Command{
	Use:   "notify-workers",
	Short: "The `NotifyWorkers` operation sends an email to one or more Workers that you specify with the Worker ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_notifyWorkersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_notifyWorkersCmd).Standalone()

		mturk_notifyWorkersCmd.Flags().String("message-text", "", "The text of the email message to send.")
		mturk_notifyWorkersCmd.Flags().String("subject", "", "The subject line of the email message to send.")
		mturk_notifyWorkersCmd.Flags().String("worker-ids", "", "A list of Worker IDs you wish to notify.")
		mturk_notifyWorkersCmd.MarkFlagRequired("message-text")
		mturk_notifyWorkersCmd.MarkFlagRequired("subject")
		mturk_notifyWorkersCmd.MarkFlagRequired("worker-ids")
	})
	mturkCmd.AddCommand(mturk_notifyWorkersCmd)
}
