package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listRoomsCmd = &cobra.Command{
	Use:   "list-rooms",
	Short: "Lists the room details for the specified Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listRoomsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_listRoomsCmd).Standalone()

		chime_listRoomsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_listRoomsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chime_listRoomsCmd.Flags().String("member-id", "", "The member ID (user ID or bot ID).")
		chime_listRoomsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		chime_listRoomsCmd.MarkFlagRequired("account-id")
	})
	chimeCmd.AddCommand(chime_listRoomsCmd)
}
