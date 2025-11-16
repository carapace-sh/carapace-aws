package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listGameServerGroupsCmd = &cobra.Command{
	Use:   "list-game-server-groups",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nLists a game server groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listGameServerGroupsCmd).Standalone()

	gamelift_listGameServerGroupsCmd.Flags().String("limit", "", "The game server groups' limit.")
	gamelift_listGameServerGroupsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	gameliftCmd.AddCommand(gamelift_listGameServerGroupsCmd)
}
