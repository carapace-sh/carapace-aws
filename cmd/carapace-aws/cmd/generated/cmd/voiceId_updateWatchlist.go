package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_updateWatchlistCmd = &cobra.Command{
	Use:   "update-watchlist",
	Short: "Updates the specified watchlist.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_updateWatchlistCmd).Standalone()

	voiceId_updateWatchlistCmd.Flags().String("description", "", "A brief description about this watchlist.")
	voiceId_updateWatchlistCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the watchlist.")
	voiceId_updateWatchlistCmd.Flags().String("name", "", "The name of the watchlist.")
	voiceId_updateWatchlistCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist to be updated.")
	voiceId_updateWatchlistCmd.MarkFlagRequired("domain-id")
	voiceId_updateWatchlistCmd.MarkFlagRequired("watchlist-id")
	voiceIdCmd.AddCommand(voiceId_updateWatchlistCmd)
}
