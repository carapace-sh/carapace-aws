package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listWatchlistsCmd = &cobra.Command{
	Use:   "list-watchlists",
	Short: "Lists all watchlists in a specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listWatchlistsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_listWatchlistsCmd).Standalone()

		voiceId_listWatchlistsCmd.Flags().String("domain-id", "", "The identifier of the domain.")
		voiceId_listWatchlistsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		voiceId_listWatchlistsCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
		voiceId_listWatchlistsCmd.MarkFlagRequired("domain-id")
	})
	voiceIdCmd.AddCommand(voiceId_listWatchlistsCmd)
}
