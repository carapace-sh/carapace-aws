package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateLexBotCmd = &cobra.Command{
	Use:   "disassociate-lex-bot",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateLexBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateLexBotCmd).Standalone()

		connect_disassociateLexBotCmd.Flags().String("bot-name", "", "The name of the Amazon Lex bot.")
		connect_disassociateLexBotCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_disassociateLexBotCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateLexBotCmd.Flags().String("lex-region", "", "The Amazon Web Services Region in which the Amazon Lex bot has been created.")
		connect_disassociateLexBotCmd.MarkFlagRequired("bot-name")
		connect_disassociateLexBotCmd.MarkFlagRequired("instance-id")
		connect_disassociateLexBotCmd.MarkFlagRequired("lex-region")
	})
	connectCmd.AddCommand(connect_disassociateLexBotCmd)
}
