package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_listStreamSessionsByAccountCmd = &cobra.Command{
	Use:   "list-stream-sessions-by-account",
	Short: "Retrieves a list of Amazon GameLift Streams stream sessions that this user account has access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_listStreamSessionsByAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_listStreamSessionsByAccountCmd).Standalone()

		gameliftstreams_listStreamSessionsByAccountCmd.Flags().String("export-files-status", "", "Filter by the exported files status.")
		gameliftstreams_listStreamSessionsByAccountCmd.Flags().String("max-results", "", "The number of results to return.")
		gameliftstreams_listStreamSessionsByAccountCmd.Flags().String("next-token", "", "The token that marks the start of the next set of results.")
		gameliftstreams_listStreamSessionsByAccountCmd.Flags().String("status", "", "Filter by the stream session status.")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_listStreamSessionsByAccountCmd)
}
