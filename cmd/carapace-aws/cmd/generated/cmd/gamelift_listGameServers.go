package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listGameServersCmd = &cobra.Command{
	Use:   "list-game-servers",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nRetrieves information on all game servers that are currently active in a specified game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listGameServersCmd).Standalone()

	gamelift_listGameServersCmd.Flags().String("game-server-group-name", "", "An identifier for the game server group to retrieve a list of game servers from.")
	gamelift_listGameServersCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_listGameServersCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_listGameServersCmd.Flags().String("sort-order", "", "Indicates how to sort the returned data based on game server registration timestamp.")
	gamelift_listGameServersCmd.MarkFlagRequired("game-server-group-name")
	gameliftCmd.AddCommand(gamelift_listGameServersCmd)
}
