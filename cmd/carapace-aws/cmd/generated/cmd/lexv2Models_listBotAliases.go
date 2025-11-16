package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotAliasesCmd = &cobra.Command{
	Use:   "list-bot-aliases",
	Short: "Gets a list of aliases for the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBotAliasesCmd).Standalone()

		lexv2Models_listBotAliasesCmd.Flags().String("bot-id", "", "The identifier of the bot to list aliases for.")
		lexv2Models_listBotAliasesCmd.Flags().String("max-results", "", "The maximum number of aliases to return in each page of results.")
		lexv2Models_listBotAliasesCmd.Flags().String("next-token", "", "If the response from the `ListBotAliases` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listBotAliasesCmd.MarkFlagRequired("bot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotAliasesCmd)
}
