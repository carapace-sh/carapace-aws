package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createRoomMembershipCmd = &cobra.Command{
	Use:   "create-room-membership",
	Short: "Adds a member to a chat room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createRoomMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_createRoomMembershipCmd).Standalone()

		chime_createRoomMembershipCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_createRoomMembershipCmd.Flags().String("member-id", "", "The Amazon Chime member ID (user ID or bot ID).")
		chime_createRoomMembershipCmd.Flags().String("role", "", "The role of the member.")
		chime_createRoomMembershipCmd.Flags().String("room-id", "", "The room ID.")
		chime_createRoomMembershipCmd.MarkFlagRequired("account-id")
		chime_createRoomMembershipCmd.MarkFlagRequired("member-id")
		chime_createRoomMembershipCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_createRoomMembershipCmd)
}
