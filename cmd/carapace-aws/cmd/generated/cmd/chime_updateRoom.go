package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateRoomCmd = &cobra.Command{
	Use:   "update-room",
	Short: "Updates room details, such as the room name, for a room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_updateRoomCmd).Standalone()

		chime_updateRoomCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_updateRoomCmd.Flags().String("name", "", "The room name.")
		chime_updateRoomCmd.Flags().String("room-id", "", "The room ID.")
		chime_updateRoomCmd.MarkFlagRequired("account-id")
		chime_updateRoomCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_updateRoomCmd)
}
