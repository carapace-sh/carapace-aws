package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_startMatchmakingCmd = &cobra.Command{
	Use:   "start-matchmaking",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUses FlexMatch to create a game match for a group of players based on custom matchmaking rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_startMatchmakingCmd).Standalone()

	gamelift_startMatchmakingCmd.Flags().String("configuration-name", "", "Name of the matchmaking configuration to use for this request.")
	gamelift_startMatchmakingCmd.Flags().String("players", "", "Information on each player to be matched.")
	gamelift_startMatchmakingCmd.Flags().String("ticket-id", "", "A unique identifier for a matchmaking ticket.")
	gamelift_startMatchmakingCmd.MarkFlagRequired("configuration-name")
	gamelift_startMatchmakingCmd.MarkFlagRequired("players")
	gameliftCmd.AddCommand(gamelift_startMatchmakingCmd)
}
