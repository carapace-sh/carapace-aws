package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_createWatchlistCmd = &cobra.Command{
	Use:   "create-watchlist",
	Short: "Creates a watchlist that fraudsters can be a part of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_createWatchlistCmd).Standalone()

	voiceId_createWatchlistCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	voiceId_createWatchlistCmd.Flags().String("description", "", "A brief description of this watchlist.")
	voiceId_createWatchlistCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the watchlist.")
	voiceId_createWatchlistCmd.Flags().String("name", "", "The name of the watchlist.")
	voiceId_createWatchlistCmd.MarkFlagRequired("domain-id")
	voiceId_createWatchlistCmd.MarkFlagRequired("name")
	voiceIdCmd.AddCommand(voiceId_createWatchlistCmd)
}
