package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateLexBotCmd = &cobra.Command{
	Use:   "associate-lex-bot",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateLexBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_associateLexBotCmd).Standalone()

		connect_associateLexBotCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_associateLexBotCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_associateLexBotCmd.Flags().String("lex-bot", "", "The Amazon Lex bot to associate with the instance.")
		connect_associateLexBotCmd.MarkFlagRequired("instance-id")
		connect_associateLexBotCmd.MarkFlagRequired("lex-bot")
	})
	connectCmd.AddCommand(connect_associateLexBotCmd)
}
