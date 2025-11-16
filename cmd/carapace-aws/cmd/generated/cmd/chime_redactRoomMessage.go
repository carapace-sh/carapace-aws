package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_redactRoomMessageCmd = &cobra.Command{
	Use:   "redact-room-message",
	Short: "Redacts the specified message from the specified Amazon Chime channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_redactRoomMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_redactRoomMessageCmd).Standalone()

		chime_redactRoomMessageCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_redactRoomMessageCmd.Flags().String("message-id", "", "The message ID.")
		chime_redactRoomMessageCmd.Flags().String("room-id", "", "The room ID.")
		chime_redactRoomMessageCmd.MarkFlagRequired("account-id")
		chime_redactRoomMessageCmd.MarkFlagRequired("message-id")
		chime_redactRoomMessageCmd.MarkFlagRequired("room-id")
	})
	chimeCmd.AddCommand(chime_redactRoomMessageCmd)
}
