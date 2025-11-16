package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueMaxContactsCmd = &cobra.Command{
	Use:   "update-queue-max-contacts",
	Short: "Updates the maximum number of contacts allowed in a queue before it is considered full.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueMaxContactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQueueMaxContactsCmd).Standalone()

		connect_updateQueueMaxContactsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQueueMaxContactsCmd.Flags().String("max-contacts", "", "The maximum number of contacts that can be in the queue before it is considered full.")
		connect_updateQueueMaxContactsCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_updateQueueMaxContactsCmd.MarkFlagRequired("instance-id")
		connect_updateQueueMaxContactsCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_updateQueueMaxContactsCmd)
}
