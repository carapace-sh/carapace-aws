package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeQueueCmd = &cobra.Command{
	Use:   "describe-queue",
	Short: "Describes the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeQueueCmd).Standalone()

		connect_describeQueueCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeQueueCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_describeQueueCmd.MarkFlagRequired("instance-id")
		connect_describeQueueCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_describeQueueCmd)
}
