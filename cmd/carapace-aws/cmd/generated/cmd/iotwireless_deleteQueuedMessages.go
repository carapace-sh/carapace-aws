package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteQueuedMessagesCmd = &cobra.Command{
	Use:   "delete-queued-messages",
	Short: "Remove queued messages from the downlink queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteQueuedMessagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteQueuedMessagesCmd).Standalone()

		iotwireless_deleteQueuedMessagesCmd.Flags().String("id", "", "The ID of a given wireless device for which downlink messages will be deleted.")
		iotwireless_deleteQueuedMessagesCmd.Flags().String("message-id", "", "If message ID is `\"*\"`, it cleares the entire downlink queue for a given device, specified by the wireless device ID.")
		iotwireless_deleteQueuedMessagesCmd.Flags().String("wireless-device-type", "", "The wireless device type, which can be either Sidewalk or LoRaWAN.")
		iotwireless_deleteQueuedMessagesCmd.MarkFlagRequired("id")
		iotwireless_deleteQueuedMessagesCmd.MarkFlagRequired("message-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteQueuedMessagesCmd)
}
