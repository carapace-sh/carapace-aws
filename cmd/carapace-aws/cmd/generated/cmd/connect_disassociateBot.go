package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateBotCmd = &cobra.Command{
	Use:   "disassociate-bot",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateBotCmd).Standalone()

		connect_disassociateBotCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_disassociateBotCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateBotCmd.Flags().String("lex-bot", "", "")
		connect_disassociateBotCmd.Flags().String("lex-v2-bot", "", "The Amazon Lex V2 bot to disassociate from the instance.")
		connect_disassociateBotCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_disassociateBotCmd)
}
