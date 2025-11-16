package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueOutboundCallerConfigCmd = &cobra.Command{
	Use:   "update-queue-outbound-caller-config",
	Short: "Updates the outbound caller ID name, number, and outbound whisper flow for a specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueOutboundCallerConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQueueOutboundCallerConfigCmd).Standalone()

		connect_updateQueueOutboundCallerConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQueueOutboundCallerConfigCmd.Flags().String("outbound-caller-config", "", "The outbound caller ID name, number, and outbound whisper flow.")
		connect_updateQueueOutboundCallerConfigCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_updateQueueOutboundCallerConfigCmd.MarkFlagRequired("instance-id")
		connect_updateQueueOutboundCallerConfigCmd.MarkFlagRequired("outbound-caller-config")
		connect_updateQueueOutboundCallerConfigCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_updateQueueOutboundCallerConfigCmd)
}
