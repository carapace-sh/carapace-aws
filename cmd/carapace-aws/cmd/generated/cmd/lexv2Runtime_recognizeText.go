package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_recognizeTextCmd = &cobra.Command{
	Use:   "recognize-text",
	Short: "Sends user input to Amazon Lex V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_recognizeTextCmd).Standalone()

	lexv2Runtime_recognizeTextCmd.Flags().String("bot-alias-id", "", "The alias identifier in use for the bot that processes the request.")
	lexv2Runtime_recognizeTextCmd.Flags().String("bot-id", "", "The identifier of the bot that processes the request.")
	lexv2Runtime_recognizeTextCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
	lexv2Runtime_recognizeTextCmd.Flags().String("request-attributes", "", "Request-specific information passed between the client application and Amazon Lex V2")
	lexv2Runtime_recognizeTextCmd.Flags().String("session-id", "", "The identifier of the user session that is having the conversation.")
	lexv2Runtime_recognizeTextCmd.Flags().String("session-state", "", "The current state of the dialog between the user and the bot.")
	lexv2Runtime_recognizeTextCmd.Flags().String("text", "", "The text that the user entered.")
	lexv2Runtime_recognizeTextCmd.MarkFlagRequired("bot-alias-id")
	lexv2Runtime_recognizeTextCmd.MarkFlagRequired("bot-id")
	lexv2Runtime_recognizeTextCmd.MarkFlagRequired("locale-id")
	lexv2Runtime_recognizeTextCmd.MarkFlagRequired("session-id")
	lexv2Runtime_recognizeTextCmd.MarkFlagRequired("text")
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_recognizeTextCmd)
}
