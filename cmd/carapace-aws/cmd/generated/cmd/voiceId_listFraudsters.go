package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listFraudstersCmd = &cobra.Command{
	Use:   "list-fraudsters",
	Short: "Lists all fraudsters in a specified watchlist or domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listFraudstersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_listFraudstersCmd).Standalone()

		voiceId_listFraudstersCmd.Flags().String("domain-id", "", "The identifier of the domain.")
		voiceId_listFraudstersCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		voiceId_listFraudstersCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
		voiceId_listFraudstersCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist.")
		voiceId_listFraudstersCmd.MarkFlagRequired("domain-id")
	})
	voiceIdCmd.AddCommand(voiceId_listFraudstersCmd)
}
