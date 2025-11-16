package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeWatchlistCmd = &cobra.Command{
	Use:   "describe-watchlist",
	Short: "Describes the specified watchlist.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeWatchlistCmd).Standalone()

	voiceId_describeWatchlistCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the watchlist.")
	voiceId_describeWatchlistCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist that you are describing.")
	voiceId_describeWatchlistCmd.MarkFlagRequired("domain-id")
	voiceId_describeWatchlistCmd.MarkFlagRequired("watchlist-id")
	voiceIdCmd.AddCommand(voiceId_describeWatchlistCmd)
}
