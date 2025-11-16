package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateRoomMembershipCmd = &cobra.Command{
	Use:   "update-room-membership",
	Short: "Updates room membership details, such as the member role, for a room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateRoomMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_updateRoomMembershipCmd).Standalone()

		chime_updateRoomMembershipCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_updateRoomMembershipCmd.Flags().String("member-id", "", "The member ID.")
		chime_updateRoomMembershipCmd.Flags().String("role", "", "The role of the member.")
		chime_updateRoomMembershipCmd.Flags().String("room-id", "", "The room ID.")
		chime_updateRoomMembershipCmd.MarkFlagRequired("account-id")
		chime_updateRoomMembershipCmd.MarkFlagRequired("member-id")
		chime_updateRoomMembershipCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_updateRoomMembershipCmd)
}
