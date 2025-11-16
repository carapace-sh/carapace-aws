package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listImportsCmd = &cobra.Command{
	Use:   "list-imports",
	Short: "Lists the imports for a bot, bot locale, or custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listImportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listImportsCmd).Standalone()

		lexv2Models_listImportsCmd.Flags().String("bot-id", "", "The unique identifier that Amazon Lex assigned to the bot.")
		lexv2Models_listImportsCmd.Flags().String("bot-version", "", "The version of the bot to list imports for.")
		lexv2Models_listImportsCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the bots in the response to only those that match the filter specification.")
		lexv2Models_listImportsCmd.Flags().String("locale-id", "", "Specifies the locale that should be present in the list.")
		lexv2Models_listImportsCmd.Flags().String("max-results", "", "The maximum number of imports to return in each page of results.")
		lexv2Models_listImportsCmd.Flags().String("next-token", "", "If the response from the `ListImports` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listImportsCmd.Flags().String("sort-by", "", "Determines the field that the list of imports is sorted by.")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listImportsCmd)
}
