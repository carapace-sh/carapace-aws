package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntime_postTextCmd = &cobra.Command{
	Use:   "post-text",
	Short: "Sends user input to Amazon Lex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntime_postTextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexRuntime_postTextCmd).Standalone()

		lexRuntime_postTextCmd.Flags().String("active-contexts", "", "A list of contexts active for the request.")
		lexRuntime_postTextCmd.Flags().String("bot-alias", "", "The alias of the Amazon Lex bot.")
		lexRuntime_postTextCmd.Flags().String("bot-name", "", "The name of the Amazon Lex bot.")
		lexRuntime_postTextCmd.Flags().String("input-text", "", "The text that the user entered (Amazon Lex interprets this text).")
		lexRuntime_postTextCmd.Flags().String("request-attributes", "", "Request-specific information passed between Amazon Lex and a client application.")
		lexRuntime_postTextCmd.Flags().String("session-attributes", "", "Application-specific information passed between Amazon Lex and a client application.")
		lexRuntime_postTextCmd.Flags().String("user-id", "", "The ID of the client application user.")
		lexRuntime_postTextCmd.MarkFlagRequired("bot-alias")
		lexRuntime_postTextCmd.MarkFlagRequired("bot-name")
		lexRuntime_postTextCmd.MarkFlagRequired("input-text")
		lexRuntime_postTextCmd.MarkFlagRequired("user-id")
	})
	lexRuntimeCmd.AddCommand(lexRuntime_postTextCmd)
}
