package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listRecommendedIntentsCmd = &cobra.Command{
	Use:   "list-recommended-intents",
	Short: "Gets a list of recommended intents provided by the bot recommendation that you can use in your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listRecommendedIntentsCmd).Standalone()

	lexv2Models_listRecommendedIntentsCmd.Flags().String("bot-id", "", "The unique identifier of the bot associated with the recommended intents.")
	lexv2Models_listRecommendedIntentsCmd.Flags().String("bot-recommendation-id", "", "The identifier of the bot recommendation that contains the recommended intents.")
	lexv2Models_listRecommendedIntentsCmd.Flags().String("bot-version", "", "The version of the bot that contains the recommended intents.")
	lexv2Models_listRecommendedIntentsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the recommended intents.")
	lexv2Models_listRecommendedIntentsCmd.Flags().String("max-results", "", "The maximum number of bot recommendations to return in each page of results.")
	lexv2Models_listRecommendedIntentsCmd.Flags().String("next-token", "", "If the response from the ListRecommendedIntents operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
	lexv2Models_listRecommendedIntentsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listRecommendedIntentsCmd.MarkFlagRequired("bot-recommendation-id")
	lexv2Models_listRecommendedIntentsCmd.MarkFlagRequired("bot-version")
	lexv2Models_listRecommendedIntentsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listRecommendedIntentsCmd)
}
