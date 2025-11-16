package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_deleteMessageCmd = &cobra.Command{
	Use:   "delete-message",
	Short: "Sends an event to a specific room which directs clients to delete a specific message; that is, unrender it from view and delete it from the client’s chat history.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_deleteMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_deleteMessageCmd).Standalone()

		ivschat_deleteMessageCmd.Flags().String("id", "", "ID of the message to be deleted.")
		ivschat_deleteMessageCmd.Flags().String("reason", "", "Reason for deleting the message.")
		ivschat_deleteMessageCmd.Flags().String("room-identifier", "", "Identifier of the room where the message should be deleted.")
		ivschat_deleteMessageCmd.MarkFlagRequired("id")
		ivschat_deleteMessageCmd.MarkFlagRequired("room-identifier")
	})
	ivschatCmd.AddCommand(ivschat_deleteMessageCmd)
}
