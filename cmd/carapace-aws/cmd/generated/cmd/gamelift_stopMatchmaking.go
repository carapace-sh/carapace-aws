package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_stopMatchmakingCmd = &cobra.Command{
	Use:   "stop-matchmaking",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCancels a matchmaking ticket or match backfill ticket that is currently being processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_stopMatchmakingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_stopMatchmakingCmd).Standalone()

		gamelift_stopMatchmakingCmd.Flags().String("ticket-id", "", "A unique identifier for a matchmaking ticket.")
		gamelift_stopMatchmakingCmd.MarkFlagRequired("ticket-id")
	})
	gameliftCmd.AddCommand(gamelift_stopMatchmakingCmd)
}
