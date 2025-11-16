package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listQueuedMessagesCmd = &cobra.Command{
	Use:   "list-queued-messages",
	Short: "List queued messages in the downlink queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listQueuedMessagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listQueuedMessagesCmd).Standalone()

		iotwireless_listQueuedMessagesCmd.Flags().String("id", "", "The ID of a given wireless device which the downlink message packets are being sent.")
		iotwireless_listQueuedMessagesCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listQueuedMessagesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iotwireless_listQueuedMessagesCmd.Flags().String("wireless-device-type", "", "The wireless device type, whic can be either Sidewalk or LoRaWAN.")
		iotwireless_listQueuedMessagesCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_listQueuedMessagesCmd)
}
