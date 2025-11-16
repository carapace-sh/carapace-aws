package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_updateRoomCmd = &cobra.Command{
	Use:   "update-room",
	Short: "Updates a room’s configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_updateRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_updateRoomCmd).Standalone()

		ivschat_updateRoomCmd.Flags().String("identifier", "", "Identifier of the room to be updated.")
		ivschat_updateRoomCmd.Flags().String("logging-configuration-identifiers", "", "Array of logging-configuration identifiers attached to the room.")
		ivschat_updateRoomCmd.Flags().String("maximum-message-length", "", "The maximum number of characters in a single message.")
		ivschat_updateRoomCmd.Flags().String("maximum-message-rate-per-second", "", "Maximum number of messages per second that can be sent to the room (by all clients).")
		ivschat_updateRoomCmd.Flags().String("message-review-handler", "", "Configuration information for optional review of messages.")
		ivschat_updateRoomCmd.Flags().String("name", "", "Room name.")
		ivschat_updateRoomCmd.MarkFlagRequired("identifier")
	})
	ivschatCmd.AddCommand(ivschat_updateRoomCmd)
}
