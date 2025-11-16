package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBuiltInIntentsCmd = &cobra.Command{
	Use:   "list-built-in-intents",
	Short: "Gets a list of built-in intents provided by Amazon Lex that you can use in your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBuiltInIntentsCmd).Standalone()

	lexv2Models_listBuiltInIntentsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the intents to list.")
	lexv2Models_listBuiltInIntentsCmd.Flags().String("max-results", "", "The maximum number of built-in intents to return in each page of results.")
	lexv2Models_listBuiltInIntentsCmd.Flags().String("next-token", "", "If the response from the `ListBuiltInIntents` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listBuiltInIntentsCmd.Flags().String("sort-by", "", "Specifies sorting parameters for the list of built-in intents.")
	lexv2Models_listBuiltInIntentsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listBuiltInIntentsCmd)
}
