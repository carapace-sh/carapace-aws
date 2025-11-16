package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeMatchmakingCmd = &cobra.Command{
	Use:   "describe-matchmaking",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves one or more matchmaking tickets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeMatchmakingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeMatchmakingCmd).Standalone()

		gamelift_describeMatchmakingCmd.Flags().String("ticket-ids", "", "A unique identifier for a matchmaking ticket.")
		gamelift_describeMatchmakingCmd.MarkFlagRequired("ticket-ids")
	})
	gameliftCmd.AddCommand(gamelift_describeMatchmakingCmd)
}
