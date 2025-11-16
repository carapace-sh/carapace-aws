package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_notifyRecommendationsReceivedCmd = &cobra.Command{
	Use:   "notify-recommendations-received",
	Short: "Removes the specified recommendations from the specified assistant's queue of newly available recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_notifyRecommendationsReceivedCmd).Standalone()

	wisdom_notifyRecommendationsReceivedCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_notifyRecommendationsReceivedCmd.Flags().String("recommendation-ids", "", "The identifiers of the recommendations.")
	wisdom_notifyRecommendationsReceivedCmd.Flags().String("session-id", "", "The identifier of the session.")
	wisdom_notifyRecommendationsReceivedCmd.MarkFlagRequired("assistant-id")
	wisdom_notifyRecommendationsReceivedCmd.MarkFlagRequired("recommendation-ids")
	wisdom_notifyRecommendationsReceivedCmd.MarkFlagRequired("session-id")
	wisdomCmd.AddCommand(wisdom_notifyRecommendationsReceivedCmd)
}
