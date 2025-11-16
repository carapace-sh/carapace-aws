package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_stopGameSessionPlacementCmd = &cobra.Command{
	Use:   "stop-game-session-placement",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCancels a game session placement that's in `PENDING` status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_stopGameSessionPlacementCmd).Standalone()

	gamelift_stopGameSessionPlacementCmd.Flags().String("placement-id", "", "A unique identifier for a game session placement to stop.")
	gamelift_stopGameSessionPlacementCmd.MarkFlagRequired("placement-id")
	gameliftCmd.AddCommand(gamelift_stopGameSessionPlacementCmd)
}
