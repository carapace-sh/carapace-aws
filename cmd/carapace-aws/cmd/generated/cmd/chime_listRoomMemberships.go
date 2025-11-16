package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listRoomMembershipsCmd = &cobra.Command{
	Use:   "list-room-memberships",
	Short: "Lists the membership details for the specified room in an Amazon Chime Enterprise account, such as the members' IDs, email addresses, and names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listRoomMembershipsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_listRoomMembershipsCmd).Standalone()

		chime_listRoomMembershipsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_listRoomMembershipsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chime_listRoomMembershipsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		chime_listRoomMembershipsCmd.Flags().String("room-id", "", "The room ID.")
		chime_listRoomMembershipsCmd.MarkFlagRequired("account-id")
		chime_listRoomMembershipsCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_listRoomMembershipsCmd)
}
