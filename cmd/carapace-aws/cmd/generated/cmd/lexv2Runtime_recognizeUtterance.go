package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_recognizeUtteranceCmd = &cobra.Command{
	Use:   "recognize-utterance",
	Short: "Sends user input to Amazon Lex V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_recognizeUtteranceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Runtime_recognizeUtteranceCmd).Standalone()

		lexv2Runtime_recognizeUtteranceCmd.Flags().String("bot-alias-id", "", "The alias identifier in use for the bot that should receive the request.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("bot-id", "", "The identifier of the bot that should receive the request.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("input-stream", "", "User input in PCM or Opus audio format or text format as described in the `requestContentType` parameter.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("request-attributes", "", "Request-specific information passed between the client application and Amazon Lex V2")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("request-content-type", "", "Indicates the format for audio input or that the content is text.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("response-content-type", "", "The message that Amazon Lex V2 returns in the response can be either text or speech based on the `responseContentType` value.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("session-id", "", "The identifier of the session in use.")
		lexv2Runtime_recognizeUtteranceCmd.Flags().String("session-state", "", "Sets the state of the session with the user.")
		lexv2Runtime_recognizeUtteranceCmd.MarkFlagRequired("bot-alias-id")
		lexv2Runtime_recognizeUtteranceCmd.MarkFlagRequired("bot-id")
		lexv2Runtime_recognizeUtteranceCmd.MarkFlagRequired("locale-id")
		lexv2Runtime_recognizeUtteranceCmd.MarkFlagRequired("request-content-type")
		lexv2Runtime_recognizeUtteranceCmd.MarkFlagRequired("session-id")
	})
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_recognizeUtteranceCmd)
}
