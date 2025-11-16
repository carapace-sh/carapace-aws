package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listIntentsCmd = &cobra.Command{
	Use:   "list-intents",
	Short: "Get a list of intents that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listIntentsCmd).Standalone()

	lexv2Models_listIntentsCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the intent.")
	lexv2Models_listIntentsCmd.Flags().String("bot-version", "", "The version of the bot that contains the intent.")
	lexv2Models_listIntentsCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the intents in the response to only those that match the filter specification.")
	lexv2Models_listIntentsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the intents to list.")
	lexv2Models_listIntentsCmd.Flags().String("max-results", "", "The maximum number of intents to return in each page of results.")
	lexv2Models_listIntentsCmd.Flags().String("next-token", "", "If the response from the `ListIntents` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listIntentsCmd.Flags().String("sort-by", "", "Determines the sort order for the response from the `ListIntents` operation.")
	lexv2Models_listIntentsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listIntentsCmd.MarkFlagRequired("bot-version")
	lexv2Models_listIntentsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listIntentsCmd)
}
