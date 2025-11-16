package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameServerGroupCmd = &cobra.Command{
	Use:   "describe-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nRetrieves information on a game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameServerGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameServerGroupCmd).Standalone()

		gamelift_describeGameServerGroupCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
		gamelift_describeGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
	})
	gameliftCmd.AddCommand(gamelift_describeGameServerGroupCmd)
}
