package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotAliasCmd = &cobra.Command{
	Use:   "get-bot-alias",
	Short: "Returns information about an Amazon Lex bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBotAliasCmd).Standalone()

		lexModels_getBotAliasCmd.Flags().String("bot-name", "", "The name of the bot.")
		lexModels_getBotAliasCmd.Flags().String("name", "", "The name of the bot alias.")
		lexModels_getBotAliasCmd.MarkFlagRequired("bot-name")
		lexModels_getBotAliasCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_getBotAliasCmd)
}
