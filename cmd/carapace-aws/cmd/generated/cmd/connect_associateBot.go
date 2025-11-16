package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateBotCmd = &cobra.Command{
	Use:   "associate-bot",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateBotCmd).Standalone()

	connect_associateBotCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_associateBotCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateBotCmd.Flags().String("lex-bot", "", "")
	connect_associateBotCmd.Flags().String("lex-v2-bot", "", "The Amazon Lex V2 bot to associate with the instance.")
	connect_associateBotCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_associateBotCmd)
}
