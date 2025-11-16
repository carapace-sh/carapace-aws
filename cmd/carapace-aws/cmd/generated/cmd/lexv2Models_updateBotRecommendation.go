package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateBotRecommendationCmd = &cobra.Command{
	Use:   "update-bot-recommendation",
	Short: "Updates an existing bot recommendation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateBotRecommendationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateBotRecommendationCmd).Standalone()

		lexv2Models_updateBotRecommendationCmd.Flags().String("bot-id", "", "The unique identifier of the bot containing the bot recommendation to be updated.")
		lexv2Models_updateBotRecommendationCmd.Flags().String("bot-recommendation-id", "", "The unique identifier of the bot recommendation to be updated.")
		lexv2Models_updateBotRecommendationCmd.Flags().String("bot-version", "", "The version of the bot containing the bot recommendation to be updated.")
		lexv2Models_updateBotRecommendationCmd.Flags().String("encryption-setting", "", "The object representing the passwords that will be used to encrypt the data related to the bot recommendation results, as well as the KMS key ARN used to encrypt the associated metadata.")
		lexv2Models_updateBotRecommendationCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the bot recommendation to update.")
		lexv2Models_updateBotRecommendationCmd.MarkFlagRequired("bot-id")
		lexv2Models_updateBotRecommendationCmd.MarkFlagRequired("bot-recommendation-id")
		lexv2Models_updateBotRecommendationCmd.MarkFlagRequired("bot-version")
		lexv2Models_updateBotRecommendationCmd.MarkFlagRequired("encryption-setting")
		lexv2Models_updateBotRecommendationCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateBotRecommendationCmd)
}
