package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listScriptsCmd = &cobra.Command{
	Use:   "list-scripts",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves script records for all Realtime scripts that are associated with the Amazon Web Services account in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listScriptsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listScriptsCmd).Standalone()

		gamelift_listScriptsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listScriptsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_listScriptsCmd)
}
