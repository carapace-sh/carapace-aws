package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotLocalesCmd = &cobra.Command{
	Use:   "list-bot-locales",
	Short: "Gets a list of locales for the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotLocalesCmd).Standalone()

	lexv2Models_listBotLocalesCmd.Flags().String("bot-id", "", "The identifier of the bot to list locales for.")
	lexv2Models_listBotLocalesCmd.Flags().String("bot-version", "", "The version of the bot to list locales for.")
	lexv2Models_listBotLocalesCmd.Flags().String("filters", "", "Provides the specification for a filter used to limit the response to only those locales that match the filter specification.")
	lexv2Models_listBotLocalesCmd.Flags().String("max-results", "", "The maximum number of aliases to return in each page of results.")
	lexv2Models_listBotLocalesCmd.Flags().String("next-token", "", "If the response from the `ListBotLocales` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listBotLocalesCmd.Flags().String("sort-by", "", "Specifies sorting parameters for the list of locales.")
	lexv2Models_listBotLocalesCmd.MarkFlagRequired("bot-id")
	lexv2Models_listBotLocalesCmd.MarkFlagRequired("bot-version")
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotLocalesCmd)
}
