package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createGameSessionQueueCmd = &cobra.Command{
	Use:   "create-game-session-queue",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCreates a placement queue that processes requests for new game sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createGameSessionQueueCmd).Standalone()

	gamelift_createGameSessionQueueCmd.Flags().String("custom-event-data", "", "Information to be added to all events that are related to this game session queue.")
	gamelift_createGameSessionQueueCmd.Flags().String("destinations", "", "A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.")
	gamelift_createGameSessionQueueCmd.Flags().String("filter-configuration", "", "A list of locations where a queue is allowed to place new game sessions.")
	gamelift_createGameSessionQueueCmd.Flags().String("name", "", "A descriptive label that is associated with game session queue.")
	gamelift_createGameSessionQueueCmd.Flags().String("notification-target", "", "An SNS topic ARN that is set up to receive game session placement notifications.")
	gamelift_createGameSessionQueueCmd.Flags().String("player-latency-policies", "", "A set of policies that enforce a sliding cap on player latency when processing game sessions placement requests.")
	gamelift_createGameSessionQueueCmd.Flags().String("priority-configuration", "", "Custom settings to use when prioritizing destinations and locations for game session placements.")
	gamelift_createGameSessionQueueCmd.Flags().String("tags", "", "A list of labels to assign to the new game session queue resource.")
	gamelift_createGameSessionQueueCmd.Flags().String("timeout-in-seconds", "", "The maximum time, in seconds, that a new game session placement request remains in the queue.")
	gamelift_createGameSessionQueueCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_createGameSessionQueueCmd)
}
