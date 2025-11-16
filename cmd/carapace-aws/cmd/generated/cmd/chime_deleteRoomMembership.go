package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_deleteRoomMembershipCmd = &cobra.Command{
	Use:   "delete-room-membership",
	Short: "Removes a member from a chat room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_deleteRoomMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_deleteRoomMembershipCmd).Standalone()

		chime_deleteRoomMembershipCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_deleteRoomMembershipCmd.Flags().String("member-id", "", "The member ID (user ID or bot ID).")
		chime_deleteRoomMembershipCmd.Flags().String("room-id", "", "The room ID.")
		chime_deleteRoomMembershipCmd.MarkFlagRequired("account-id")
		chime_deleteRoomMembershipCmd.MarkFlagRequired("member-id")
		chime_deleteRoomMembershipCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_deleteRoomMembershipCmd)
}
