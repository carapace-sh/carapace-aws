package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotRecommendationCmd = &cobra.Command{
	Use:   "describe-bot-recommendation",
	Short: "Provides metadata information about a bot recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotRecommendationCmd).Standalone()

	lexv2Models_describeBotRecommendationCmd.Flags().String("bot-id", "", "The unique identifier of the bot associated with the bot recommendation.")
	lexv2Models_describeBotRecommendationCmd.Flags().String("bot-recommendation-id", "", "The identifier of the bot recommendation to describe.")
	lexv2Models_describeBotRecommendationCmd.Flags().String("bot-version", "", "The version of the bot associated with the bot recommendation.")
	lexv2Models_describeBotRecommendationCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the bot recommendation to describe.")
	lexv2Models_describeBotRecommendationCmd.MarkFlagRequired("bot-id")
	lexv2Models_describeBotRecommendationCmd.MarkFlagRequired("bot-recommendation-id")
	lexv2Models_describeBotRecommendationCmd.MarkFlagRequired("bot-version")
	lexv2Models_describeBotRecommendationCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotRecommendationCmd)
}
