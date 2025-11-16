package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getRecommendationsCmd = &cobra.Command{
	Use:   "get-recommendations",
	Short: "Retrieves recommendations for the specified session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getRecommendationsCmd).Standalone()

	wisdom_getRecommendationsCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_getRecommendationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_getRecommendationsCmd.Flags().String("session-id", "", "The identifier of the session.")
	wisdom_getRecommendationsCmd.Flags().String("wait-time-seconds", "", "The duration (in seconds) for which the call waits for a recommendation to be made available before returning.")
	wisdom_getRecommendationsCmd.MarkFlagRequired("assistant-id")
	wisdom_getRecommendationsCmd.MarkFlagRequired("session-id")
	wisdomCmd.AddCommand(wisdom_getRecommendationsCmd)
}
