package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_listStreamSessionsCmd = &cobra.Command{
	Use:   "list-stream-sessions",
	Short: "Retrieves a list of Amazon GameLift Streams stream sessions that a stream group is hosting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_listStreamSessionsCmd).Standalone()

	gameliftstreams_listStreamSessionsCmd.Flags().String("export-files-status", "", "Filter by the exported files status.")
	gameliftstreams_listStreamSessionsCmd.Flags().String("identifier", "", "The unique identifier of a Amazon GameLift Streams stream group to retrieve the stream session for.")
	gameliftstreams_listStreamSessionsCmd.Flags().String("max-results", "", "The number of results to return.")
	gameliftstreams_listStreamSessionsCmd.Flags().String("next-token", "", "The token that marks the start of the next set of results.")
	gameliftstreams_listStreamSessionsCmd.Flags().String("status", "", "Filter by the stream session status.")
	gameliftstreams_listStreamSessionsCmd.MarkFlagRequired("identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_listStreamSessionsCmd)
}
