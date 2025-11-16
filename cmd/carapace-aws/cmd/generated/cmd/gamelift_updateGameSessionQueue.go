package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateGameSessionQueueCmd = &cobra.Command{
	Use:   "update-game-session-queue",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUpdates the configuration of a game session queue, which determines how the queue processes new game session requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateGameSessionQueueCmd).Standalone()

	gamelift_updateGameSessionQueueCmd.Flags().String("custom-event-data", "", "Information to be added to all events that are related to this game session queue.")
	gamelift_updateGameSessionQueueCmd.Flags().String("destinations", "", "A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.")
	gamelift_updateGameSessionQueueCmd.Flags().String("filter-configuration", "", "A list of locations where a queue is allowed to place new game sessions.")
	gamelift_updateGameSessionQueueCmd.Flags().String("name", "", "A descriptive label that is associated with game session queue.")
	gamelift_updateGameSessionQueueCmd.Flags().String("notification-target", "", "An SNS topic ARN that is set up to receive game session placement notifications.")
	gamelift_updateGameSessionQueueCmd.Flags().String("player-latency-policies", "", "A set of policies that enforce a sliding cap on player latency when processing game sessions placement requests.")
	gamelift_updateGameSessionQueueCmd.Flags().String("priority-configuration", "", "Custom settings to use when prioritizing destinations and locations for game session placements.")
	gamelift_updateGameSessionQueueCmd.Flags().String("timeout-in-seconds", "", "The maximum time, in seconds, that a new game session placement request remains in the queue.")
	gamelift_updateGameSessionQueueCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_updateGameSessionQueueCmd)
}
