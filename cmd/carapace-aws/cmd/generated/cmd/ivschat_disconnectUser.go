package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_disconnectUserCmd = &cobra.Command{
	Use:   "disconnect-user",
	Short: "Disconnects all connections using a specified user ID from a room.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_disconnectUserCmd).Standalone()

	ivschat_disconnectUserCmd.Flags().String("reason", "", "Reason for disconnecting the user.")
	ivschat_disconnectUserCmd.Flags().String("room-identifier", "", "Identifier of the room from which the user's clients should be disconnected.")
	ivschat_disconnectUserCmd.Flags().String("user-id", "", "ID of the user (connection) to disconnect from the room.")
	ivschat_disconnectUserCmd.MarkFlagRequired("room-identifier")
	ivschat_disconnectUserCmd.MarkFlagRequired("user-id")
	ivschatCmd.AddCommand(ivschat_disconnectUserCmd)
}
