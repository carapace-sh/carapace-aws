package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_putSessionCmd = &cobra.Command{
	Use:   "put-session",
	Short: "Creates a new session or modifies an existing session with an Amazon Lex V2 bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_putSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Runtime_putSessionCmd).Standalone()

		lexv2Runtime_putSessionCmd.Flags().String("bot-alias-id", "", "The alias identifier of the bot that receives the session data.")
		lexv2Runtime_putSessionCmd.Flags().String("bot-id", "", "The identifier of the bot that receives the session data.")
		lexv2Runtime_putSessionCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
		lexv2Runtime_putSessionCmd.Flags().String("messages", "", "A list of messages to send to the user.")
		lexv2Runtime_putSessionCmd.Flags().String("request-attributes", "", "Request-specific information passed between Amazon Lex V2 and the client application.")
		lexv2Runtime_putSessionCmd.Flags().String("response-content-type", "", "The message that Amazon Lex V2 returns in the response can be either text or speech depending on the value of this parameter.")
		lexv2Runtime_putSessionCmd.Flags().String("session-id", "", "The identifier of the session that receives the session data.")
		lexv2Runtime_putSessionCmd.Flags().String("session-state", "", "Sets the state of the session with the user.")
		lexv2Runtime_putSessionCmd.MarkFlagRequired("bot-alias-id")
		lexv2Runtime_putSessionCmd.MarkFlagRequired("bot-id")
		lexv2Runtime_putSessionCmd.MarkFlagRequired("locale-id")
		lexv2Runtime_putSessionCmd.MarkFlagRequired("session-id")
		lexv2Runtime_putSessionCmd.MarkFlagRequired("session-state")
	})
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_putSessionCmd)
}
