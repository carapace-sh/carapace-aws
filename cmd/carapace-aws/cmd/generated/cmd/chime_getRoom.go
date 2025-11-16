package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getRoomCmd = &cobra.Command{
	Use:   "get-room",
	Short: "Retrieves room details, such as the room name, for a room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_getRoomCmd).Standalone()

		chime_getRoomCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_getRoomCmd.Flags().String("room-id", "", "The room ID.")
		chime_getRoomCmd.MarkFlagRequired("account-id")
		chime_getRoomCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_getRoomCmd)
}
