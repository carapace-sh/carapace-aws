package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getRecommendationsCmd = &cobra.Command{
	Use:   "get-recommendations",
	Short: "This API will be discontinued starting June 1, 2024. To receive generative responses after March 1, 2024, you will need to create a new Assistant in the Amazon Connect console and integrate the Amazon Q in Connect JavaScript library (amazon-q-connectjs) into your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getRecommendationsCmd).Standalone()

		qconnect_getRecommendationsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_getRecommendationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_getRecommendationsCmd.Flags().String("next-chunk-token", "", "The token for the next set of chunks.")
		qconnect_getRecommendationsCmd.Flags().String("session-id", "", "The identifier of the session.")
		qconnect_getRecommendationsCmd.Flags().String("wait-time-seconds", "", "The duration (in seconds) for which the call waits for a recommendation to be made available before returning.")
		qconnect_getRecommendationsCmd.MarkFlagRequired("assistant-id")
		qconnect_getRecommendationsCmd.MarkFlagRequired("session-id")
	})
	qconnectCmd.AddCommand(qconnect_getRecommendationsCmd)
}
