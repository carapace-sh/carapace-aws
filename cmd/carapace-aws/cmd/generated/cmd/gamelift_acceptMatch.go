package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_acceptMatchCmd = &cobra.Command{
	Use:   "accept-match",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRegisters a player's acceptance or rejection of a proposed FlexMatch match.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_acceptMatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_acceptMatchCmd).Standalone()

		gamelift_acceptMatchCmd.Flags().String("acceptance-type", "", "Player response to the proposed match.")
		gamelift_acceptMatchCmd.Flags().String("player-ids", "", "A unique identifier for a player delivering the response.")
		gamelift_acceptMatchCmd.Flags().String("ticket-id", "", "A unique identifier for a matchmaking ticket.")
		gamelift_acceptMatchCmd.MarkFlagRequired("acceptance-type")
		gamelift_acceptMatchCmd.MarkFlagRequired("player-ids")
		gamelift_acceptMatchCmd.MarkFlagRequired("ticket-id")
	})
	gameliftCmd.AddCommand(gamelift_acceptMatchCmd)
}
