package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listFleetsCmd = &cobra.Command{
	Use:   "list-fleets",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves a collection of fleet resources in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listFleetsCmd).Standalone()

		gamelift_listFleetsCmd.Flags().String("build-id", "", "A unique identifier for the build to request fleets for.")
		gamelift_listFleetsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listFleetsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_listFleetsCmd.Flags().String("script-id", "", "A unique identifier for the Realtime script to request fleets for.")
	})
	gameliftCmd.AddCommand(gamelift_listFleetsCmd)
}
