package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_stopBotRecommendationCmd = &cobra.Command{
	Use:   "stop-bot-recommendation",
	Short: "Stop an already running Bot Recommendation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_stopBotRecommendationCmd).Standalone()

	lexv2Models_stopBotRecommendationCmd.Flags().String("bot-id", "", "The unique identifier of the bot containing the bot recommendation to be stopped.")
	lexv2Models_stopBotRecommendationCmd.Flags().String("bot-recommendation-id", "", "The unique identifier of the bot recommendation to be stopped.")
	lexv2Models_stopBotRecommendationCmd.Flags().String("bot-version", "", "The version of the bot containing the bot recommendation.")
	lexv2Models_stopBotRecommendationCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the bot recommendation to stop.")
	lexv2Models_stopBotRecommendationCmd.MarkFlagRequired("bot-id")
	lexv2Models_stopBotRecommendationCmd.MarkFlagRequired("bot-recommendation-id")
	lexv2Models_stopBotRecommendationCmd.MarkFlagRequired("bot-version")
	lexv2Models_stopBotRecommendationCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_stopBotRecommendationCmd)
}
