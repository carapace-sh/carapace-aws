package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_startBotRecommendationCmd = &cobra.Command{
	Use:   "start-bot-recommendation",
	Short: "Use this to provide your transcript data, and to start the bot recommendation process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_startBotRecommendationCmd).Standalone()

	lexv2Models_startBotRecommendationCmd.Flags().String("bot-id", "", "The unique identifier of the bot containing the bot recommendation.")
	lexv2Models_startBotRecommendationCmd.Flags().String("bot-version", "", "The version of the bot containing the bot recommendation.")
	lexv2Models_startBotRecommendationCmd.Flags().String("encryption-setting", "", "The object representing the passwords that will be used to encrypt the data related to the bot recommendation results, as well as the KMS key ARN used to encrypt the associated metadata.")
	lexv2Models_startBotRecommendationCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the bot recommendation to start.")
	lexv2Models_startBotRecommendationCmd.Flags().String("transcript-source-setting", "", "The object representing the Amazon S3 bucket containing the transcript, as well as the associated metadata.")
	lexv2Models_startBotRecommendationCmd.MarkFlagRequired("bot-id")
	lexv2Models_startBotRecommendationCmd.MarkFlagRequired("bot-version")
	lexv2Models_startBotRecommendationCmd.MarkFlagRequired("locale-id")
	lexv2Models_startBotRecommendationCmd.MarkFlagRequired("transcript-source-setting")
	lexv2ModelsCmd.AddCommand(lexv2Models_startBotRecommendationCmd)
}
