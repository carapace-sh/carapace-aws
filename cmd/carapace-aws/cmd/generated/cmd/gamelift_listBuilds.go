package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listBuildsCmd = &cobra.Command{
	Use:   "list-builds",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves build resources for all builds associated with the Amazon Web Services account in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listBuildsCmd).Standalone()

	gamelift_listBuildsCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_listBuildsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_listBuildsCmd.Flags().String("status", "", "Build status to filter results by.")
	gameliftCmd.AddCommand(gamelift_listBuildsCmd)
}
