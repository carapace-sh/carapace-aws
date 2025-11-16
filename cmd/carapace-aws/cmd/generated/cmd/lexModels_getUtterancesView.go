package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getUtterancesViewCmd = &cobra.Command{
	Use:   "get-utterances-view",
	Short: "Use the `GetUtterancesView` operation to get information about the utterances that your users have made to your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getUtterancesViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getUtterancesViewCmd).Standalone()

		lexModels_getUtterancesViewCmd.Flags().String("bot-name", "", "The name of the bot for which utterance information should be returned.")
		lexModels_getUtterancesViewCmd.Flags().String("bot-versions", "", "An array of bot versions for which utterance information should be returned.")
		lexModels_getUtterancesViewCmd.Flags().String("status-type", "", "To return utterances that were recognized and handled, use `Detected`.")
		lexModels_getUtterancesViewCmd.MarkFlagRequired("bot-name")
		lexModels_getUtterancesViewCmd.MarkFlagRequired("bot-versions")
		lexModels_getUtterancesViewCmd.MarkFlagRequired("status-type")
	})
	lexModelsCmd.AddCommand(lexModels_getUtterancesViewCmd)
}
