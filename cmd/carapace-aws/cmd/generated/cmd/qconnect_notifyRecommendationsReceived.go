package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_notifyRecommendationsReceivedCmd = &cobra.Command{
	Use:   "notify-recommendations-received",
	Short: "Removes the specified recommendations from the specified assistant's queue of newly available recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_notifyRecommendationsReceivedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_notifyRecommendationsReceivedCmd).Standalone()

		qconnect_notifyRecommendationsReceivedCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_notifyRecommendationsReceivedCmd.Flags().String("recommendation-ids", "", "The identifiers of the recommendations.")
		qconnect_notifyRecommendationsReceivedCmd.Flags().String("session-id", "", "The identifier of the session.")
		qconnect_notifyRecommendationsReceivedCmd.MarkFlagRequired("assistant-id")
		qconnect_notifyRecommendationsReceivedCmd.MarkFlagRequired("recommendation-ids")
		qconnect_notifyRecommendationsReceivedCmd.MarkFlagRequired("session-id")
	})
	qconnectCmd.AddCommand(qconnect_notifyRecommendationsReceivedCmd)
}
