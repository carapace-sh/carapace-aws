package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotRecommendationsCmd = &cobra.Command{
	Use:   "list-bot-recommendations",
	Short: "Get a list of bot recommendations that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotRecommendationsCmd).Standalone()

	lexv2Models_listBotRecommendationsCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the bot recommendation list.")
	lexv2Models_listBotRecommendationsCmd.Flags().String("bot-version", "", "The version of the bot that contains the bot recommendation list.")
	lexv2Models_listBotRecommendationsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the bot recommendation list.")
	lexv2Models_listBotRecommendationsCmd.Flags().String("max-results", "", "The maximum number of bot recommendations to return in each page of results.")
	lexv2Models_listBotRecommendationsCmd.Flags().String("next-token", "", "If the response from the ListBotRecommendation operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
	lexv2Models_listBotRecommendationsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listBotRecommendationsCmd.MarkFlagRequired("bot-version")
	lexv2Models_listBotRecommendationsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotRecommendationsCmd)
}
