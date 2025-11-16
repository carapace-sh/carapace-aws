package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_sendEventCmd = &cobra.Command{
	Use:   "send-event",
	Short: "Sends an event to a room.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_sendEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_sendEventCmd).Standalone()

		ivschat_sendEventCmd.Flags().String("attributes", "", "Application-defined metadata to attach to the event sent to clients.")
		ivschat_sendEventCmd.Flags().String("event-name", "", "Application-defined name of the event to send to clients.")
		ivschat_sendEventCmd.Flags().String("room-identifier", "", "Identifier of the room to which the event will be sent.")
		ivschat_sendEventCmd.MarkFlagRequired("event-name")
		ivschat_sendEventCmd.MarkFlagRequired("room-identifier")
	})
	ivschatCmd.AddCommand(ivschat_sendEventCmd)
}
