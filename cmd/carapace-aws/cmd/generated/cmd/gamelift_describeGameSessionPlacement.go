package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameSessionPlacementCmd = &cobra.Command{
	Use:   "describe-game-session-placement",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves information, including current status, about a game session placement request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameSessionPlacementCmd).Standalone()

	gamelift_describeGameSessionPlacementCmd.Flags().String("placement-id", "", "A unique identifier for a game session placement to retrieve.")
	gamelift_describeGameSessionPlacementCmd.MarkFlagRequired("placement-id")
	gameliftCmd.AddCommand(gamelift_describeGameSessionPlacementCmd)
}
