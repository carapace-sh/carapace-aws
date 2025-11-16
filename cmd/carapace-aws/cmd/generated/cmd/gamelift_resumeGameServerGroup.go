package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_resumeGameServerGroupCmd = &cobra.Command{
	Use:   "resume-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nReinstates activity on a game server group after it has been suspended.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_resumeGameServerGroupCmd).Standalone()

	gamelift_resumeGameServerGroupCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
	gamelift_resumeGameServerGroupCmd.Flags().String("resume-actions", "", "The activity to resume for this game server group.")
	gamelift_resumeGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
	gamelift_resumeGameServerGroupCmd.MarkFlagRequired("resume-actions")
	gameliftCmd.AddCommand(gamelift_resumeGameServerGroupCmd)
}
