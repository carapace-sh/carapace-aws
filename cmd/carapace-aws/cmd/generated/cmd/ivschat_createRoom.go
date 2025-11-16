package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_createRoomCmd = &cobra.Command{
	Use:   "create-room",
	Short: "Creates a room that allows clients to connect and pass messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_createRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_createRoomCmd).Standalone()

		ivschat_createRoomCmd.Flags().String("logging-configuration-identifiers", "", "Array of logging-configuration identifiers attached to the room.")
		ivschat_createRoomCmd.Flags().String("maximum-message-length", "", "Maximum number of characters in a single message.")
		ivschat_createRoomCmd.Flags().String("maximum-message-rate-per-second", "", "Maximum number of messages per second that can be sent to the room (by all clients).")
		ivschat_createRoomCmd.Flags().String("message-review-handler", "", "Configuration information for optional review of messages.")
		ivschat_createRoomCmd.Flags().String("name", "", "Room name.")
		ivschat_createRoomCmd.Flags().String("tags", "", "Tags to attach to the resource.")
	})
	ivschatCmd.AddCommand(ivschat_createRoomCmd)
}
