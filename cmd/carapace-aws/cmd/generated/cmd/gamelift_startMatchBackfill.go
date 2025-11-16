package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_startMatchBackfillCmd = &cobra.Command{
	Use:   "start-match-backfill",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nFinds new players to fill open slots in currently running game sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_startMatchBackfillCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_startMatchBackfillCmd).Standalone()

		gamelift_startMatchBackfillCmd.Flags().String("configuration-name", "", "Name of the matchmaker to use for this request.")
		gamelift_startMatchBackfillCmd.Flags().String("game-session-arn", "", "A unique identifier for the game session.")
		gamelift_startMatchBackfillCmd.Flags().String("players", "", "Match information on all players that are currently assigned to the game session.")
		gamelift_startMatchBackfillCmd.Flags().String("ticket-id", "", "A unique identifier for a matchmaking ticket.")
		gamelift_startMatchBackfillCmd.MarkFlagRequired("configuration-name")
		gamelift_startMatchBackfillCmd.MarkFlagRequired("players")
	})
	gameliftCmd.AddCommand(gamelift_startMatchBackfillCmd)
}
