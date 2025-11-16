package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntime_postContentCmd = &cobra.Command{
	Use:   "post-content",
	Short: "Sends user input (text or speech) to Amazon Lex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntime_postContentCmd).Standalone()

	lexRuntime_postContentCmd.Flags().String("accept", "", "You pass this value as the `Accept` HTTP header.")
	lexRuntime_postContentCmd.Flags().String("active-contexts", "", "A list of contexts active for the request.")
	lexRuntime_postContentCmd.Flags().String("bot-alias", "", "Alias of the Amazon Lex bot.")
	lexRuntime_postContentCmd.Flags().String("bot-name", "", "Name of the Amazon Lex bot.")
	lexRuntime_postContentCmd.Flags().String("content-type", "", "You pass this value as the `Content-Type` HTTP header.")
	lexRuntime_postContentCmd.Flags().String("input-stream", "", "User input in PCM or Opus audio format or text format as described in the `Content-Type` HTTP header.")
	lexRuntime_postContentCmd.Flags().String("request-attributes", "", "You pass this value as the `x-amz-lex-request-attributes` HTTP header.")
	lexRuntime_postContentCmd.Flags().String("session-attributes", "", "You pass this value as the `x-amz-lex-session-attributes` HTTP header.")
	lexRuntime_postContentCmd.Flags().String("user-id", "", "The ID of the client application user.")
	lexRuntime_postContentCmd.MarkFlagRequired("bot-alias")
	lexRuntime_postContentCmd.MarkFlagRequired("bot-name")
	lexRuntime_postContentCmd.MarkFlagRequired("content-type")
	lexRuntime_postContentCmd.MarkFlagRequired("input-stream")
	lexRuntime_postContentCmd.MarkFlagRequired("user-id")
	lexRuntimeCmd.AddCommand(lexRuntime_postContentCmd)
}
