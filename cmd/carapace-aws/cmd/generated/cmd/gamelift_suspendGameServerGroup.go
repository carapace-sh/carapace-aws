package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_suspendGameServerGroupCmd = &cobra.Command{
	Use:   "suspend-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nTemporarily stops activity on a game server group without terminating instances or the game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_suspendGameServerGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_suspendGameServerGroupCmd).Standalone()

		gamelift_suspendGameServerGroupCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
		gamelift_suspendGameServerGroupCmd.Flags().String("suspend-actions", "", "The activity to suspend for this game server group.")
		gamelift_suspendGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
		gamelift_suspendGameServerGroupCmd.MarkFlagRequired("suspend-actions")
	})
	gameliftCmd.AddCommand(gamelift_suspendGameServerGroupCmd)
}
