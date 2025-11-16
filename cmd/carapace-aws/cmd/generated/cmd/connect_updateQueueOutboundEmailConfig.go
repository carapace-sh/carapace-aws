package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueOutboundEmailConfigCmd = &cobra.Command{
	Use:   "update-queue-outbound-email-config",
	Short: "Updates the outbound email address Id for a specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueOutboundEmailConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQueueOutboundEmailConfigCmd).Standalone()

		connect_updateQueueOutboundEmailConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQueueOutboundEmailConfigCmd.Flags().String("outbound-email-config", "", "The outbound email address ID for a specified queue.")
		connect_updateQueueOutboundEmailConfigCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_updateQueueOutboundEmailConfigCmd.MarkFlagRequired("instance-id")
		connect_updateQueueOutboundEmailConfigCmd.MarkFlagRequired("outbound-email-config")
		connect_updateQueueOutboundEmailConfigCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_updateQueueOutboundEmailConfigCmd)
}
