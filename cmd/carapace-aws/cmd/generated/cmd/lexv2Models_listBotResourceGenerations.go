package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotResourceGenerationsCmd = &cobra.Command{
	Use:   "list-bot-resource-generations",
	Short: "Lists the generation requests made for a bot locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotResourceGenerationsCmd).Standalone()

	lexv2Models_listBotResourceGenerationsCmd.Flags().String("bot-id", "", "The unique identifier of the bot whose generation requests you want to view.")
	lexv2Models_listBotResourceGenerationsCmd.Flags().String("bot-version", "", "The version of the bot whose generation requests you want to view.")
	lexv2Models_listBotResourceGenerationsCmd.Flags().String("locale-id", "", "The locale of the bot whose generation requests you want to view.")
	lexv2Models_listBotResourceGenerationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	lexv2Models_listBotResourceGenerationsCmd.Flags().String("next-token", "", "If the total number of results is greater than the number specified in the `maxResults`, the response returns a token in the `nextToken` field.")
	lexv2Models_listBotResourceGenerationsCmd.Flags().String("sort-by", "", "An object containing information about the attribute and the method by which to sort the results")
	lexv2Models_listBotResourceGenerationsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listBotResourceGenerationsCmd.MarkFlagRequired("bot-version")
	lexv2Models_listBotResourceGenerationsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotResourceGenerationsCmd)
}
