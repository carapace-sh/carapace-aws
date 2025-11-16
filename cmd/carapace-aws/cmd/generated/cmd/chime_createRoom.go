package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createRoomCmd = &cobra.Command{
	Use:   "create-room",
	Short: "Creates a chat room for the specified Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createRoomCmd).Standalone()

	chime_createRoomCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_createRoomCmd.Flags().String("client-request-token", "", "The idempotency token for the request.")
	chime_createRoomCmd.Flags().String("name", "", "The room name.")
	chime_createRoomCmd.MarkFlagRequired("account-id")
	chime_createRoomCmd.MarkFlagRequired("name")
	chimeCmd.AddCommand(chime_createRoomCmd)
}
