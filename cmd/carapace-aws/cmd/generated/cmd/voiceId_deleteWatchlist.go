package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_deleteWatchlistCmd = &cobra.Command{
	Use:   "delete-watchlist",
	Short: "Deletes the specified watchlist from Voice ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_deleteWatchlistCmd).Standalone()

	voiceId_deleteWatchlistCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the watchlist.")
	voiceId_deleteWatchlistCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist to be deleted.")
	voiceId_deleteWatchlistCmd.MarkFlagRequired("domain-id")
	voiceId_deleteWatchlistCmd.MarkFlagRequired("watchlist-id")
	voiceIdCmd.AddCommand(voiceId_deleteWatchlistCmd)
}
