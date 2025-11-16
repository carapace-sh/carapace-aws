package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotAliasesCmd = &cobra.Command{
	Use:   "get-bot-aliases",
	Short: "Returns a list of aliases for a specified Amazon Lex bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBotAliasesCmd).Standalone()

		lexModels_getBotAliasesCmd.Flags().String("bot-name", "", "The name of the bot.")
		lexModels_getBotAliasesCmd.Flags().String("max-results", "", "The maximum number of aliases to return in the response.")
		lexModels_getBotAliasesCmd.Flags().String("name-contains", "", "Substring to match in bot alias names.")
		lexModels_getBotAliasesCmd.Flags().String("next-token", "", "A pagination token for fetching the next page of aliases.")
		lexModels_getBotAliasesCmd.MarkFlagRequired("bot-name")
	})
	lexModelsCmd.AddCommand(lexModels_getBotAliasesCmd)
}
