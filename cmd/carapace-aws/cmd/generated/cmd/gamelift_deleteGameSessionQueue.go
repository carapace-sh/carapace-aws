package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteGameSessionQueueCmd = &cobra.Command{
	Use:   "delete-game-session-queue",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nDeletes a game session queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteGameSessionQueueCmd).Standalone()

	gamelift_deleteGameSessionQueueCmd.Flags().String("name", "", "A descriptive label that is associated with game session queue.")
	gamelift_deleteGameSessionQueueCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_deleteGameSessionQueueCmd)
}
