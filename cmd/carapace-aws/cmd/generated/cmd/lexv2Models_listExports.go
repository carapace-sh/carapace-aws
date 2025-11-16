package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: "Lists the exports for a bot, bot locale, or custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listExportsCmd).Standalone()

		lexv2Models_listExportsCmd.Flags().String("bot-id", "", "The unique identifier that Amazon Lex assigned to the bot.")
		lexv2Models_listExportsCmd.Flags().String("bot-version", "", "The version of the bot to list exports for.")
		lexv2Models_listExportsCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the exports in the response to only those that match the filter specification.")
		lexv2Models_listExportsCmd.Flags().String("locale-id", "", "Specifies the resources that should be exported.")
		lexv2Models_listExportsCmd.Flags().String("max-results", "", "The maximum number of exports to return in each page of results.")
		lexv2Models_listExportsCmd.Flags().String("next-token", "", "If the response from the `ListExports` operation contains more results that specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listExportsCmd.Flags().String("sort-by", "", "Determines the field that the list of exports is sorted by.")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listExportsCmd)
}
