package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueNameCmd = &cobra.Command{
	Use:   "update-queue-name",
	Short: "Updates the name and description of a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQueueNameCmd).Standalone()

		connect_updateQueueNameCmd.Flags().String("description", "", "The description of the queue.")
		connect_updateQueueNameCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQueueNameCmd.Flags().String("name", "", "The name of the queue.")
		connect_updateQueueNameCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_updateQueueNameCmd.MarkFlagRequired("instance-id")
		connect_updateQueueNameCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_updateQueueNameCmd)
}
