package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_startConversationCmd = &cobra.Command{
	Use:   "start-conversation",
	Short: "Starts an HTTP/2 bidirectional event stream that enables you to send audio, text, or DTMF input in real time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_startConversationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Runtime_startConversationCmd).Standalone()

		lexv2Runtime_startConversationCmd.Flags().String("bot-alias-id", "", "The alias identifier in use for the bot that processes the request.")
		lexv2Runtime_startConversationCmd.Flags().String("bot-id", "", "The identifier of the bot to process the request.")
		lexv2Runtime_startConversationCmd.Flags().String("conversation-mode", "", "The conversation type that you are using the Amazon Lex V2.")
		lexv2Runtime_startConversationCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
		lexv2Runtime_startConversationCmd.Flags().String("request-event-stream", "", "Represents the stream of events to Amazon Lex V2 from your application.")
		lexv2Runtime_startConversationCmd.Flags().String("session-id", "", "The identifier of the user session that is having the conversation.")
		lexv2Runtime_startConversationCmd.MarkFlagRequired("bot-alias-id")
		lexv2Runtime_startConversationCmd.MarkFlagRequired("bot-id")
		lexv2Runtime_startConversationCmd.MarkFlagRequired("locale-id")
		lexv2Runtime_startConversationCmd.MarkFlagRequired("request-event-stream")
		lexv2Runtime_startConversationCmd.MarkFlagRequired("session-id")
	})
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_startConversationCmd)
}
