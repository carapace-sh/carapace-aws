package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteUtterancesCmd = &cobra.Command{
	Use:   "delete-utterances",
	Short: "Deletes stored utterances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteUtterancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteUtterancesCmd).Standalone()

		lexModels_deleteUtterancesCmd.Flags().String("bot-name", "", "The name of the bot that stored the utterances.")
		lexModels_deleteUtterancesCmd.Flags().String("user-id", "", "The unique identifier for the user that made the utterances.")
		lexModels_deleteUtterancesCmd.MarkFlagRequired("bot-name")
		lexModels_deleteUtterancesCmd.MarkFlagRequired("user-id")
	})
	lexModelsCmd.AddCommand(lexModels_deleteUtterancesCmd)
}
