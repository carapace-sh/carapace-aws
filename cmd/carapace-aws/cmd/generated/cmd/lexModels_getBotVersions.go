package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotVersionsCmd = &cobra.Command{
	Use:   "get-bot-versions",
	Short: "Gets information about all of the versions of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBotVersionsCmd).Standalone()

		lexModels_getBotVersionsCmd.Flags().String("max-results", "", "The maximum number of bot versions to return in the response.")
		lexModels_getBotVersionsCmd.Flags().String("name", "", "The name of the bot for which versions should be returned.")
		lexModels_getBotVersionsCmd.Flags().String("next-token", "", "A pagination token for fetching the next page of bot versions.")
		lexModels_getBotVersionsCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_getBotVersionsCmd)
}
