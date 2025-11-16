package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_createChatTokenCmd = &cobra.Command{
	Use:   "create-chat-token",
	Short: "Creates an encrypted token that is used by a chat participant to establish an individual WebSocket chat connection to a room.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_createChatTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_createChatTokenCmd).Standalone()

		ivschat_createChatTokenCmd.Flags().String("attributes", "", "Application-provided attributes to encode into the token and attach to a chat session.")
		ivschat_createChatTokenCmd.Flags().String("capabilities", "", "Set of capabilities that the user is allowed to perform in the room.")
		ivschat_createChatTokenCmd.Flags().String("room-identifier", "", "Identifier of the room that the client is trying to access.")
		ivschat_createChatTokenCmd.Flags().String("session-duration-in-minutes", "", "Session duration (in minutes), after which the session expires.")
		ivschat_createChatTokenCmd.Flags().String("user-id", "", "Application-provided ID that uniquely identifies the user associated with this token.")
		ivschat_createChatTokenCmd.MarkFlagRequired("room-identifier")
		ivschat_createChatTokenCmd.MarkFlagRequired("user-id")
	})
	ivschatCmd.AddCommand(ivschat_createChatTokenCmd)
}
