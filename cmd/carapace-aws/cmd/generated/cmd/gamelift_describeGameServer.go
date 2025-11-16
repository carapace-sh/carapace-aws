package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameServerCmd = &cobra.Command{
	Use:   "describe-game-server",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nRetrieves information for a registered game server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameServerCmd).Standalone()

		gamelift_describeGameServerCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group where the game server is running.")
		gamelift_describeGameServerCmd.Flags().String("game-server-id", "", "A custom string that uniquely identifies the game server information to be retrieved.")
		gamelift_describeGameServerCmd.MarkFlagRequired("game-server-group-name")
		gamelift_describeGameServerCmd.MarkFlagRequired("game-server-id")
	})
	gameliftCmd.AddCommand(gamelift_describeGameServerCmd)
}
