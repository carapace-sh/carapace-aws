package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteUtterancesCmd = &cobra.Command{
	Use:   "delete-utterances",
	Short: "Deletes stored utterances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteUtterancesCmd).Standalone()

	lexv2Models_deleteUtterancesCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the utterances.")
	lexv2Models_deleteUtterancesCmd.Flags().String("locale-id", "", "The identifier of the language and locale where the utterances were collected.")
	lexv2Models_deleteUtterancesCmd.Flags().String("session-id", "", "The unique identifier of the session with the user.")
	lexv2Models_deleteUtterancesCmd.MarkFlagRequired("bot-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteUtterancesCmd)
}
