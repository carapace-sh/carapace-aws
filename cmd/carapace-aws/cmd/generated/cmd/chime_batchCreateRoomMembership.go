package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchCreateRoomMembershipCmd = &cobra.Command{
	Use:   "batch-create-room-membership",
	Short: "Adds up to 50 members to a chat room in an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchCreateRoomMembershipCmd).Standalone()

	chime_batchCreateRoomMembershipCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_batchCreateRoomMembershipCmd.Flags().String("membership-item-list", "", "The list of membership items.")
	chime_batchCreateRoomMembershipCmd.Flags().String("room-id", "", "The room ID.")
	chime_batchCreateRoomMembershipCmd.MarkFlagRequired("account-id")
	chime_batchCreateRoomMembershipCmd.MarkFlagRequired("membership-item-list")
	chime_batchCreateRoomMembershipCmd.MarkFlagRequired("room-id")
	chimeCmd.AddCommand(chime_batchCreateRoomMembershipCmd)
}
