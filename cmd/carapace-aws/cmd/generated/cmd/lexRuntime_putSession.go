package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntime_putSessionCmd = &cobra.Command{
	Use:   "put-session",
	Short: "Creates a new session or modifies an existing session with an Amazon Lex bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntime_putSessionCmd).Standalone()

	lexRuntime_putSessionCmd.Flags().String("accept", "", "The message that Amazon Lex returns in the response can be either text or speech based depending on the value of this field.")
	lexRuntime_putSessionCmd.Flags().String("active-contexts", "", "A list of contexts active for the request.")
	lexRuntime_putSessionCmd.Flags().String("bot-alias", "", "The alias in use for the bot that contains the session data.")
	lexRuntime_putSessionCmd.Flags().String("bot-name", "", "The name of the bot that contains the session data.")
	lexRuntime_putSessionCmd.Flags().String("dialog-action", "", "Sets the next action that the bot should take to fulfill the conversation.")
	lexRuntime_putSessionCmd.Flags().String("recent-intent-summary-view", "", "A summary of the recent intents for the bot.")
	lexRuntime_putSessionCmd.Flags().String("session-attributes", "", "Map of key/value pairs representing the session-specific context information.")
	lexRuntime_putSessionCmd.Flags().String("user-id", "", "The ID of the client application user.")
	lexRuntime_putSessionCmd.MarkFlagRequired("bot-alias")
	lexRuntime_putSessionCmd.MarkFlagRequired("bot-name")
	lexRuntime_putSessionCmd.MarkFlagRequired("user-id")
	lexRuntimeCmd.AddCommand(lexRuntime_putSessionCmd)
}
