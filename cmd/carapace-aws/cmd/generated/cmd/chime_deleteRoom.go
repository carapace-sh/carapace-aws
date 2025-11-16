package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_deleteRoomCmd = &cobra.Command{
	Use:   "delete-room",
	Short: "Deletes a chat room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_deleteRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_deleteRoomCmd).Standalone()

		chime_deleteRoomCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_deleteRoomCmd.Flags().String("room-id", "", "The chat room ID.")
		chime_deleteRoomCmd.MarkFlagRequired("account-id")
		chime_deleteRoomCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_deleteRoomCmd)
}
